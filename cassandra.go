// +build !windows

package metrics

import (
	"fmt"
	"time"
)

// Output each metric in the given registry to cassandra periodically using
// the given syslogger.
func Cassandra(r Registry, d time.Duration, w *syslog.Writer) {
	for {
		r.Each(func(name string, i interface{}) {
			switch metric := i.(type) {
			case Counter:
				w.Info(fmt.Sprintf("counter %s: count: %d", name, metric.Count()))
			case Gauge:
				w.Info(fmt.Sprintf("gauge %s: value: %d", name, metric.Value()))
		})
		time.Sleep(d)
	}
}
