package collector
import(
	"net"
        "time"
	"github.com/prometheus/client_golang/prometheus"
)
type yig_9000Metrics struct {
    yig_9000Desc *prometheus.Desc
}

func (c *yig_9000Metrics) tcp_function() (flag int) {
	_,err := net.DialTimeout("tcp","192.168.2.128:9000",time.Second*10)
	if err!=nil {
            return 0
        } else{
    	    return 1
	}
}

func init() {
	registerCollector("yig_9000", defaultEnabled, Newyig_9000Metrics)
}
func Newyig_9000Metrics()  (Collector, error) {
	return &yig_9000Metrics{
        yig_9000Desc: prometheus.NewDesc(
        "tcp_yig_9000_status",
        "tcp_yig_9000_monitor",
        []string{"itemId","host"},
        nil,
        ),
     }, nil
}
func (c *yig_9000Metrics) Update(ch chan<- prometheus.Metric) error{
	flag:=c.tcp_function()
    ch <- prometheus.MustNewConstMetric(
        c.yig_9000Desc,
        prometheus.CounterValue,
        float64(flag),
        "yig_9000",
        "192.168.2.128",
    )
    return nil
}

