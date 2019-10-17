package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/go-redis/redis"
)
type {{.Type}}Metrics struct {
    {{.Type}}Desc *prometheus.Desc
}

func (c *{{.Type}}Metrics) Cache_function() (flag string){
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
	flag:=c.Cache_function()
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

