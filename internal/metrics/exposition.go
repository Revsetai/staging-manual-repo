package metrics

import (
	"bufio"
	"io"
	"math"
	"strconv"
	"strings"
)

// ContentType is the media type of the text exposition format.
const ContentType = "text/plain; version=0.0.4; charset=utf-8"

// Encoder renders registry snapshots in the Prometheus text format.
type Encoder struct {
	// Namespace is prefixed to every metric name, if set.
	Namespace string
	// Help maps a metric name to its HELP line.
	Help map[string]string
}

// NewEncoder builds an encoder for the given namespace.
func NewEncoder(namespace string) *Encoder {
	return &Encoder{Namespace: namespace, Help: make(map[string]string)}
}

// Describe records the HELP text for a metric name and returns the encoder,
// so a batch of declarations can be chained at construction.
func (e *Encoder) Describe(name, help string) *Encoder {
	if e.Help == nil {
		e.Help = make(map[string]string)
	}
	e.Help[name] = help
	return e
}

// qualify applies the namespace to a bare metric name.
func (e *Encoder) qualify(name string) string {
	if e.Namespace == "" {
		return name
	}
	return e.Namespace + "_" + name
}

// Encode writes every sample to w, grouping by metric name so that each family
// carries exactly one HELP and one TYPE line.
//
// Output is streamed through a bufio.Writer rather than accumulated in a
// strings.Builder: a scrape of a few thousand series used to allocate the
// whole page twice, once to build it and once to write it.
func (e *Encoder) Encode(w io.Writer, samples []Sample) error {
	// Snapshot already sorts by identity, so a family's samples are
	// contiguous and the run can be emitted in a single pass.
	bw := bufio.NewWriter(w)

	for start := 0; start < len(samples); {
		end := start + 1
		for end < len(samples) && samples[end].Name == samples[start].Name {
			end++
		}
		family := samples[start:end]
		start = end

		qualified := e.qualify(family[0].Name)
		if help, ok := e.Help[family[0].Name]; ok {
			bw.WriteString("# HELP ")
			bw.WriteString(qualified)
			bw.WriteByte(' ')
			bw.WriteString(helpEscaper.Replace(help))
			bw.WriteByte('\n')
		}
		bw.WriteString("# TYPE ")
		bw.WriteString(qualified)
		bw.WriteByte(' ')
		bw.WriteString(family[0].Kind.String())
		bw.WriteByte('\n')

		for _, sample := range family {
			bw.WriteString(qualified)
			writeLabels(bw, sample.Labels)
			bw.WriteByte(' ')
			bw.WriteString(formatValue(sample.Value))
			bw.WriteByte('\n')
		}
	}

	return bw.Flush()
}

// writeLabels renders a label set directly into bw, in name order.
func writeLabels(bw *bufio.Writer, labels Labels) {
	if len(labels) == 0 {
		return
	}

	bw.WriteByte('{')
	for i, name := range labels.Names() {
		if i > 0 {
			bw.WriteByte(',')
		}
		bw.WriteString(name)
		bw.WriteString(`="`)
		bw.WriteString(labelEscaper.Replace(labels[name]))
		bw.WriteByte('"')
	}
	bw.WriteByte('}')
}

func formatValue(v float64) string {
	switch {
	case math.IsNaN(v):
		return "NaN"
	case math.IsInf(v, 1):
		return "+Inf"
	case math.IsInf(v, -1):
		return "-Inf"
	default:
		return strconv.FormatFloat(v, 'g', -1, 64)
	}
}

// Replacers are built once: strings.NewReplacer compiles a trie, and doing
// that per label value made escaping the dominant cost of a large scrape.
var (
	labelEscaper = strings.NewReplacer(`\`, `\\`, `"`, `\"`, "\n", `\n`)
	helpEscaper  = strings.NewReplacer(`\`, `\\`, "\n", `\n`)
)

// EncodeRegistry is a convenience wrapper around Encode.
func (e *Encoder) EncodeRegistry(w io.Writer, r *Registry) error {
	return e.Encode(w, r.Snapshot())
}
