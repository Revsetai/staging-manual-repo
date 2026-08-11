// Package metrics provides the small, dependency-free instrument set ingestd
// uses to describe its own behaviour.
package metrics

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
)

// Labels is an unordered set of dimensions attached to a metric.
type Labels map[string]string

// key renders labels into a stable identity string.
func (l Labels) key() string {
	if len(l) == 0 {
		return ""
	}

	var b strings.Builder
	for i, name := range l.Names() {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(name)
		b.WriteByte('=')
		b.WriteString(l[name])
	}
	return b.String()
}

// Names reports every label name, in sorted order.
func (l Labels) Names() []string {
	names := make([]string, 0, len(l))
	for name := range l {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Kind distinguishes the instrument types the registry can hold.
type Kind int

const (
	// KindCounter is a monotonically increasing total.
	KindCounter Kind = iota
	// KindGauge is a value that moves in both directions.
	KindGauge
	// KindHistogram is a bucketed distribution.
	KindHistogram
	// KindInfo is a constant series whose labels carry the information; the
	// value is always 1. Build stamps and feature toggles live here rather
	// than being smuggled into a gauge nobody can read.
	KindInfo
)

// String renders the kind for export.
func (k Kind) String() string {
	switch k {
	case KindCounter:
		return "counter"
	case KindGauge:
		return "gauge"
	case KindHistogram:
		return "histogram"
	case KindInfo:
		return "gauge"
	default:
		return "untyped"
	}
}

// Counter is a monotonically increasing total.
type Counter struct {
	name   string
	labels Labels
	value  atomic.Uint64
}

// Inc adds one to the counter.
func (c *Counter) Inc() { c.Add(1) }

// Add increases the counter by delta.
func (c *Counter) Add(delta uint64) { c.value.Add(delta) }

// Value reads the current total.
func (c *Counter) Value() uint64 { return c.value.Load() }

// Name reports the counter's metric name.
func (c *Counter) Name() string { return c.name }

// Gauge is a value that can move in both directions.
//
// The value is held as the IEEE-754 bit pattern in a Uint64 so that reads and
// writes stay lock-free; Add compare-and-swaps rather than locking, which
// matters because the queue-depth gauge is written on every enqueue.
type Gauge struct {
	name   string
	labels Labels
	bits   atomic.Uint64
}

// Set replaces the gauge's value.
func (g *Gauge) Set(v float64) {
	g.bits.Store(math.Float64bits(v))
}

// Add moves the gauge by delta.
func (g *Gauge) Add(delta float64) {
	for {
		old := g.bits.Load()
		next := math.Float64bits(math.Float64frombits(old) + delta)
		if g.bits.CompareAndSwap(old, next) {
			return
		}
	}
}

// Value reads the gauge.
func (g *Gauge) Value() float64 {
	return math.Float64frombits(g.bits.Load())
}

// Name reports the gauge's metric name.
func (g *Gauge) Name() string { return g.name }

// Info is a constant series describing the process.
type Info struct {
	name   string
	labels Labels
}

// Name reports the info series' metric name.
func (i *Info) Name() string { return i.name }

// Registry owns every instrument in the process.
type Registry struct {
	mu       sync.RWMutex
	counters map[string]*Counter
	gauges   map[string]*Gauge
	infos    map[string]*Info
}

// NewRegistry builds an empty registry.
func NewRegistry() *Registry {
	return &Registry{
		counters: make(map[string]*Counter),
		gauges:   make(map[string]*Gauge),
		infos:    make(map[string]*Info),
	}
}

func identity(name string, labels Labels) string {
	if k := labels.key(); k != "" {
		return name + "{" + k + "}"
	}
	return name
}

// lookup returns the instrument registered under id, creating it with make
// if this is its first use.
//
// The read lock is taken first because the overwhelmingly common case is a
// hot path re-fetching an instrument it has already registered; only the
// first call per identity pays for the write lock.
func lookup[T any](r *Registry, into map[string]*T, id string, make func() *T) *T {
	r.mu.RLock()
	existing, ok := into[id]
	r.mu.RUnlock()
	if ok {
		return existing
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	if existing, ok := into[id]; ok {
		return existing
	}
	created := make()
	into[id] = created
	return created
}

// Counter returns the counter for name/labels, creating it on first use.
func (r *Registry) Counter(name string, labels Labels) *Counter {
	id := identity(name, labels)
	return lookup(r, r.counters, id, func() *Counter {
		return &Counter{name: name, labels: labels}
	})
}

// Gauge returns the gauge for name/labels, creating it on first use.
func (r *Registry) Gauge(name string, labels Labels) *Gauge {
	id := identity(name, labels)
	return lookup(r, r.gauges, id, func() *Gauge {
		return &Gauge{name: name, labels: labels}
	})
}

// Info registers a constant series carrying labels and the value 1.
func (r *Registry) Info(name string, labels Labels) *Info {
	id := identity(name, labels)
	return lookup(r, r.infos, id, func() *Info {
		return &Info{name: name, labels: labels}
	})
}

// Sample is a point-in-time reading of one instrument.
type Sample struct {
	// ID is the rendered name{labels} identity, carried so that callers
	// sorting or grouping samples do not rebuild it per comparison.
	ID     string
	Name   string
	Kind   Kind
	Labels Labels
	Value  float64
}

// Snapshot returns every instrument's current value, sorted by identity.
func (r *Registry) Snapshot() []Sample {
	r.mu.RLock()
	samples := make([]Sample, 0, len(r.counters)+len(r.gauges)+len(r.infos))
	for id, c := range r.counters {
		samples = append(samples, Sample{
			ID:     id,
			Name:   c.name,
			Kind:   KindCounter,
			Labels: c.labels,
			Value:  float64(c.Value()),
		})
	}
	for id, g := range r.gauges {
		samples = append(samples, Sample{
			ID:     id,
			Name:   g.name,
			Kind:   KindGauge,
			Labels: g.labels,
			Value:  g.Value(),
		})
	}
	for id, i := range r.infos {
		samples = append(samples, Sample{
			ID:     id,
			Name:   i.name,
			Kind:   KindInfo,
			Labels: i.labels,
			Value:  1,
		})
	}
	r.mu.RUnlock()

	sort.Slice(samples, func(i, j int) bool { return samples[i].ID < samples[j].ID })
	return samples
}

// String renders a sample the way a debug endpoint would.
func (s Sample) String() string {
	return fmt.Sprintf("%s %s %g", s.ID, s.Kind, s.Value)
}
