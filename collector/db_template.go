package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)
type {{.Type}}Metrics struct {
    {{.Type}}Desc *prometheus.Desc
}

func (c *{{.Type}}Metrics) DB_function() (flag string) {
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
	registerCollector("{{.Type}}", defaultEnabled, New{{.Type}}Metrics)
}
func New{{.Type}}Metrics()  (Collector, error) {
	return &{{.Type}}Metrics{
		{{.Type}}Desc: prometheus.NewDesc(
			"{{.FqName}}",
			"{{.FqName}}_monitor",
			[]string{"{{.VariableLabels}}"},
			nil,
		),
	}, nil
}
func (c *{{.Type}}Metrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.DB_function()
	if flag=="down"{
		value=0
	}
	ch <- prometheus.MustNewConstMetric(
		c.{{.Type}}Desc,
		prometheus.CounterValue,
		float64(value),
		"{{.LabelValues}}",
	)
	return nil
}


