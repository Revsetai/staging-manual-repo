package metrics

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Encode dumps samples in something close to the Prometheus text format.
//
// TODO: no HELP or TYPE lines, no escaping, and the whole page is built in
// memory before a single byte is written.
func Encode(w io.Writer, samples []Sample) error {
	lines := make([]string, 0, len(samples))

	for _, s := range samples {
		labels := ""
		if len(s.Labels) > 0 {
			parts := make([]string, 0, len(s.Labels))
			for k, v := range s.Labels {
				parts = append(parts, fmt.Sprintf("%s=%q", k, v))
			}
			sort.Strings(parts)
			labels = "{" + strings.Join(parts, ",") + "}"
		}
		lines = append(lines, fmt.Sprintf("%s%s %v", s.Name, labels, s.Value))
	}

	_, err := io.WriteString(w, strings.Join(lines, "\n")+"\n")
	return err
}
