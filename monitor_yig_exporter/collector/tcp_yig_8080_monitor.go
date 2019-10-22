package collector
import(
	"net"
        "time"
	"github.com/prometheus/client_golang/prometheus"
)
type yig_8080Metrics struct {
    yig_8080Desc *prometheus.Desc
}

func (c *yig_8080Metrics) tcp_function() (flag int) {
	_,err := net.DialTimeout("tcp","192.168.2.128:8080",time.Second*10)
	if err!=nil {
            return 0
        } else{
    	    return 1
	}
}

func init() {
	registerCollector("yig_8080", defaultEnabled, Newyig_8080Metrics)
}
func Newyig_8080Metrics()  (Collector, error) {
	return &yig_8080Metrics{
        yig_8080Desc: prometheus.NewDesc(
        "tcp_yig_8080_status",
        "tcp_yig_8080_monitor",
        []string{"itemId","host"},
        nil,
        ),
     }, nil
}
func (c *yig_8080Metrics) Update(ch chan<- prometheus.Metric) error{
	flag:=c.tcp_function()
    ch <- prometheus.MustNewConstMetric(
        c.yig_8080Desc,
        prometheus.CounterValue,
        float64(flag),
        "yig_8080",
        "192.168.2.128",
    )
    return nil
}

