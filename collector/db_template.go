package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)
type {{.Name}}Metrics struct {
    {{.Name}}Desc *prometheus.Desc
}

func (c *{{.Name}}Metrics) DB_function() (flag string) {
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
	registerCollector("{{.Name}}", defaultEnabled, New{{.Name}}Metrics)
}
func New{{.Name}}Metrics()  (Collector, error) {
	return &{{.Name}}Metrics{
		{{.Name}}Desc: prometheus.NewDesc(
			"{{.FqName}}",
			"{{.FqName}}_monitor",
			[]string{"{{.VariableLabels}}"},
			nil,
		),
	}, nil
}
func (c *{{.Name}}Metrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.DB_function()
	if flag=="down"{
		value=0
	}
	ch <- prometheus.MustNewConstMetric(
		c.{{.Name}}Desc,
		prometheus.CounterValue,
		float64(value),
		"{{.LabelValues}}",
	)
	return nil
}


