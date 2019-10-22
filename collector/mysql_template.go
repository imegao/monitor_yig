package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)
type {{.ItemId}}Metrics struct {
    {{.ItemId}}Desc *prometheus.Desc
}

func (c *{{.ItemId}}Metrics) DB_function() (flag string) {
	db, err := sql.Open("mysql", "{{.DataSourceName}}")
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
	registerCollector("{{.ItemId}}", defaultEnabled, New{{.ItemId}}Metrics)
}
func New{{.ItemId}}Metrics()  (Collector, error) {
	return &{{.ItemId}}Metrics{
		{{.ItemId}}Desc: prometheus.NewDesc(
			"mysql_{{.ItemId}}_status",
			"mysql_{{.ItemId}}_monitor",
			[]string{"itemId","host"},
			nil,
		),
	}, nil
}
func (c *{{.ItemId}}Metrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.DB_function()
	if flag=="down"{
		value=0
	}
	ch <- prometheus.MustNewConstMetric(
		c.{{.ItemId}}Desc,
		prometheus.CounterValue,
		float64(value),
		"{{.ItemId}}",
                "{{.Host}}",
	)
	return nil
}


