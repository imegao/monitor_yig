package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/go-redis/redis"
)
type {{.Name}}Metrics struct {
    {{.Name}}Desc *prometheus.Desc
}

func (c *{{.Name}}Metrics) Cache_function() (flag string){
	client := redis.NewClient(&redis.Options{
		Addr:     "{{.Addr}}",
		Password: "{{.Password}}",
		DB:       0,
	})

	defer client.Close()

	_, err := client.Ping().Result()
	if err != nil {
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
	flag:=c.Cache_function()
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

