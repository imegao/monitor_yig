package collector
import(
	"net"
        "time"
	"github.com/prometheus/client_golang/prometheus"
)
type tcp1Metrics struct {
    tcp1Desc *prometheus.Desc
}

func (c *tcp1Metrics) tcp_function() (flag int) {
	_,err := net.DialTimeout("tcp","192.168.2.128:9888",time.Second*10)
	if err!=nil {
            return 0
        } else{
    	    return 1
	}
}

func init() {
	registerCollector("tcp1", defaultEnabled, Newtcp1Metrics)
}
func Newtcp1Metrics()  (Collector, error) {
	return &tcp1Metrics{
        tcp1Desc: prometheus.NewDesc(
        "tcp1_status",
        "tcp1_status_monitor",
        []string{"status"},
        nil,
        ),
     }, nil
}
func (c *tcp1Metrics) Update(ch chan<- prometheus.Metric) error{
	flag:=c.tcp_function()
    ch <- prometheus.MustNewConstMetric(
        c.tcp1Desc,
        prometheus.CounterValue,
        float64(flag),
        "tcp",
    )
    return nil
}




