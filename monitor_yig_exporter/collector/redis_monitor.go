package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/go-redis/redis"
)
type redisMetrics struct {
    redisDesc *prometheus.Desc
}

func (c *redisMetrics) Cache_function() (flag string){
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.2.128:6379",
		Password: "hehehehe",
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
	registerCollector("redis", defaultEnabled, NewredisMetrics)
}
func NewredisMetrics()  (Collector, error) {
	return &redisMetrics{
		     redisDesc: prometheus.NewDesc(
		     	"redis_status",
			    "redis_status_monitor",
			    []string{"status"},
			    nil,
		),
	}, nil
}

func (c *redisMetrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.Cache_function()
	if flag=="down"{
		value=0
	}
	ch <- prometheus.MustNewConstMetric(
		c.redisDesc,
		prometheus.CounterValue,
		float64(value),
		"redis",
	)
	return nil
}

