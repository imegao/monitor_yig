package collector
import(
	"net"
        "time"
	"github.com/prometheus/client_golang/prometheus"
)
type {{.Type}}Metrics struct {
    {{.Type}}Desc *prometheus.Desc
}

func (c *{{.Type}}Metrics) tcp_function() (flag int) {
	_,err := net.DialTimeout("tcp","{{.Addr}}",time.Second*10)
	if err!=nil {
            return 0
        } else{
    	    return 1
	}
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
	flag:=c.tcp_function()
    ch <- prometheus.MustNewConstMetric(
        c.{{.Type}}Desc,
        prometheus.CounterValue,
        float64(flag),
        "{{.LabelValues}}",
    )
    return nil
}




