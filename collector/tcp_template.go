package collector
import(
	"net"
        "time"
	"github.com/prometheus/client_golang/prometheus"
)
type {{.Name}}Metrics struct {
    {{.Name}}Desc *prometheus.Desc
}

func (c *{{.Name}}Metrics) tcp_function() (flag int) {
	_,err := net.DialTimeout("tcp","{{.Addr}}",time.Second*10)
	if err!=nil {
            return 0
        } else{
    	    return 1
	}
}

func init() {
	registerCollector("{{.Name}}", defaultEnabled, New{{.Name}}Metrics)
}
func New{{.Name}}Metrics()  (Collector, error) {
	return &{{.Name}}Metrics{
        {{.Name}}Desc: prometheus.NewDesc(
        "{{.FqName}}",
        "{{.FqName}}_monitor",
        []string{"{{.VariableLabels}}"},
        nil,
        ),
     }, nil
}
func (c *{{.Name}}Metrics) Update(ch chan<- prometheus.Metric) error{
	flag:=c.tcp_function()
    ch <- prometheus.MustNewConstMetric(
        c.{{.Name}}Desc,
        prometheus.CounterValue,
        float64(flag),
        "{{.LabelValues}}",
    )
    return nil
}




