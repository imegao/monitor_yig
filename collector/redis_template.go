package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/go-redis/redis"
)
type {{.ItemId}}Metrics struct {
    {{.ItemId}}Desc *prometheus.Desc
}

func (c *{{.ItemId}}Metrics) Cache_function() (flag string){
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
	registerCollector("{{.ItemId}}", defaultEnabled, New{{.ItemId}}Metrics)
}
func New{{.ItemId}}Metrics()  (Collector, error) {
	return &{{.ItemId}}Metrics{
		     {{.ItemId}}Desc: prometheus.NewDesc(
		     	"redis_{{.ItemId}}_status",
			"redis_{{.ItemId}}_monitor",
			[]string{"itemId","host"},
			nil,
		),
	}, nil
}

func (c *{{.ItemId}}Metrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.Cache_function()
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

