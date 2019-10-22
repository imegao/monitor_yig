package collector
import(
	"net"
        "time"
	"github.com/prometheus/client_golang/prometheus"
)
type {{.ItemId}}Metrics struct {
    {{.ItemId}}Desc *prometheus.Desc
}

func (c *{{.ItemId}}Metrics) tcp_function() (flag int) {
	_,err := net.DialTimeout("tcp","{{.Addr}}",time.Second*10)
	if err!=nil {
            return 0
        } else{
    	    return 1
	}
}

func init() {
	registerCollector("{{.ItemId}}", defaultEnabled, New{{.ItemId}}Metrics)
}
func New{{.ItemId}}Metrics()  (Collector, error) {
	return &{{.ItemId}}Metrics{
        {{.ItemId}}Desc: prometheus.NewDesc(
        "tcp_{{.ItemId}}_status",
        "tcp_{{.ItemId}}_monitor",
        []string{"itemId","host"},
        nil,
        ),
     }, nil
}
func (c *{{.ItemId}}Metrics) Update(ch chan<- prometheus.Metric) error{
	flag:=c.tcp_function()
    ch <- prometheus.MustNewConstMetric(
        c.{{.ItemId}}Desc,
        prometheus.CounterValue,
        float64(flag),
        "{{.ItemId}}",
        "{{.Host}}",
    )
    return nil
}

