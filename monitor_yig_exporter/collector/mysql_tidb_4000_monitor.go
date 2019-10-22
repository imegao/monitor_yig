package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)
type tidb_4000Metrics struct {
    tidb_4000Desc *prometheus.Desc
}

func (c *tidb_4000Metrics) DB_function() (flag string) {
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
	registerCollector("tidb_4000", defaultEnabled, Newtidb_4000Metrics)
}
func Newtidb_4000Metrics()  (Collector, error) {
	return &tidb_4000Metrics{
		tidb_4000Desc: prometheus.NewDesc(
			"mysql_tidb_4000_status",
			"mysql_tidb_4000_monitor",
			[]string{"itemId","host"},
			nil,
		),
	}, nil
}
func (c *tidb_4000Metrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.DB_function()
	if flag=="down"{
		value=0
	}
	ch <- prometheus.MustNewConstMetric(
		c.tidb_4000Desc,
		prometheus.CounterValue,
		float64(value),
		"tidb_4000",
                "192.168.2.128",
	)
	return nil
}


