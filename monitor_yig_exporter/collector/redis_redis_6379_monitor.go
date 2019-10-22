package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/go-redis/redis"
)
type redis_6379Metrics struct {
    redis_6379Desc *prometheus.Desc
}

func (c *redis_6379Metrics) Cache_function() (flag string){
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
	registerCollector("redis_6379", defaultEnabled, Newredis_6379Metrics)
}
func Newredis_6379Metrics()  (Collector, error) {
	return &redis_6379Metrics{
		     redis_6379Desc: prometheus.NewDesc(
		     	"redis_redis_6379_status",
			"redis_redis_6379_monitor",
			[]string{"itemId","host"},
			nil,
		),
	}, nil
}

func (c *redis_6379Metrics) Update(ch chan<- prometheus.Metric) error{
	var value=1
	flag:=c.Cache_function()
	if flag=="down"{
		value=0
	}
	ch <- prometheus.MustNewConstMetric(
		c.redis_6379Desc,
		prometheus.CounterValue,
		float64(value),
		"redis_6379",
                "192.168.2.128",
	)
	return nil
}

