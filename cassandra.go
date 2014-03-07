package metrics

import (
	"fmt"
	"time"
	"github.com/gocql/gocql"
)

// Output each metric in the given registry to cassandra periodically by creating new session
// the given cassandrasyslogger.
func Cassandra(r Registry, d time.Duration, cassandra_cluster *gocql.ClusterConfig, table string, server string) {
	for {
		session, _ := cluster.CreateSession()
		r.Each(func(name string, i interface{}) {
			switch metric := i.(type) {
			case Counter:
				if err := session.Query(`INSERT INTO metrics2(server,metric, time, v) VALUES
				  (?, ?, ?, ?) USING TTL 1209600`,
				        server, name, int(time.Now().Unix()), float32(metric.Count())).Exec() ; err != nil {
				  panic(err)
				}
				// c.Clear() but this would reset all counters, need option on counter for this
			case Gauge:
				if err := session.Query(`INSERT INTO metrics2(server,metric, time, v) VALUES
				  (?, ?, ?, ?) USING TTL 1209600`,
				        server, name, int(time.Now().Unix()), float32(metric.Value())).Exec() ; err != nil {
				  panic(err)
				}
		})
		session.Close()
		time.Sleep(d)
	}
}
