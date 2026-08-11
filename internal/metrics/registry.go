// Package metrics keeps a few counters so the daemon can describe itself.
package metrics

import (
	"sort"
	"strings"
	"sync"
)

// Labels are the dimensions attached to a metric.
type Labels map[string]string

// Kind says what sort of instrument a sample came from.
type Kind int

const (
	KindCounter Kind = iota
	KindGauge
)

// String renders the kind.
func (k Kind) String() string {
	if k == KindGauge {
		return "gauge"
	}
	return "counter"
}

// Sample is one instrument's current value.
type Sample struct {
	Name   string
	Kind   Kind
	Labels Labels
	Value  float64
}

// Registry holds every metric in the process.
//
// TODO: everything is a float64 behind one mutex, so a counter and a gauge are
// indistinguishable once they are in here, and every increment on the hot path
// contends with every other one.
type Registry struct {
	mu     sync.Mutex
	values map[string]float64
	kinds  map[string]Kind
	labels map[string]Labels
	names  map[string]string
}

// NewRegistry builds an empty registry.
func NewRegistry() *Registry {
	return &Registry{
		values: make(map[string]float64),
		kinds:  make(map[string]Kind),
		labels: make(map[string]Labels),
		names:  make(map[string]string),
	}
}

func key(name string, labels Labels) string {
	if len(labels) == 0 {
		return name
	}
	parts := make([]string, 0, len(labels))
	for k, v := range labels {
		parts = append(parts, k+"="+v)
	}
	sort.Strings(parts)
	return name + "{" + strings.Join(parts, ",") + "}"
}

func (r *Registry) touch(name string, labels Labels, kind Kind) string {
	id := key(name, labels)
	r.names[id] = name
	r.labels[id] = labels
	r.kinds[id] = kind
	return id
}

// Inc adds one to a counter.
func (r *Registry) Inc(name string, labels Labels) {
	r.Add(name, labels, 1)
}

// Add increases a counter by delta.
func (r *Registry) Add(name string, labels Labels, delta float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := r.touch(name, labels, KindCounter)
	r.values[id] += delta
}

// Set replaces a gauge's value.
func (r *Registry) Set(name string, labels Labels, value float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := r.touch(name, labels, KindGauge)
	r.values[id] = value
}

// Value reads one instrument.
func (r *Registry) Value(name string, labels Labels) float64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.values[key(name, labels)]
}

// Snapshot returns every instrument, sorted by identity.
func (r *Registry) Snapshot() []Sample {
	r.mu.Lock()
	defer r.mu.Unlock()

	ids := make([]string, 0, len(r.values))
	for id := range r.values {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	samples := make([]Sample, 0, len(ids))
	for _, id := range ids {
		samples = append(samples, Sample{
			Name:   r.names[id],
			Kind:   r.kinds[id],
			Labels: r.labels[id],
			Value:  r.values[id],
		})
	}
	return samples
}
