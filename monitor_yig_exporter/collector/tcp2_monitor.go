package collector
import(
	"net"
        "time"
	"github.com/prometheus/client_golang/prometheus"
)
type tcp2Metrics struct {
    tcp2Desc *prometheus.Desc
}

func (c *tcp2Metrics) tcp_function() (flag int) {
	_,err := net.DialTimeout("tcp","192.168.2.128:9887",time.Second*10)
	if err!=nil {
            return 0
        } else{
    	    return 1
	}
}

func init() {
	registerCollector("tcp2", defaultEnabled, Newtcp2Metrics)
}
func Newtcp2Metrics()  (Collector, error) {
	return &tcp2Metrics{
        tcp2Desc: prometheus.NewDesc(
        "tcp2_status",
        "tcp2_status_monitor",
        []string{"status"},
        nil,
        ),
     }, nil
}
func (c *tcp2Metrics) Update(ch chan<- prometheus.Metric) error{
	flag:=c.tcp_function()
    ch <- prometheus.MustNewConstMetric(
        c.tcp2Desc,
        prometheus.CounterValue,
        float64(flag),
        "tcp",
    )
    return nil
}




