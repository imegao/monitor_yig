package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)
type tidbMetrics struct {
    tidbDesc *prometheus.Desc
}

func (c *tidbMetrics) DB_function() (flag string) {
	db, err := sql.Open("mysql", "root:@tcp(192.168.2.128:4000)/yig?charset=utf8")
	if err !=nil{
		fmt.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return "down"
	}
	return "up"
}

func init() {
	registerCollector("tidb", defaultEnabled, NewtidbMetrics)
}
func NewtidbMetrics()  (Collector, error) {
	return &tidbMetrics{
		tidbDesc: prometheus.NewDesc(
			"tidb_status",
			"tidb_status_monitor",
			[]string{"status"},
			nil,
		),
	}, nil
}
func (c *tidbMetrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.DB_function()
	if flag=="down"{
		value=0
	}
	ch <- prometheus.MustNewConstMetric(
		c.tidbDesc,
		prometheus.CounterValue,
		float64(value),
		"tidb",
	)
	return nil
}


