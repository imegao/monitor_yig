package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type {{.Type}}Metrics struct {
    {{.Type}}Desc *prometheus.Desc
}

func (c *{{.Type}}Metrics) Http_function() (http_status int) {
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		MaxIdleConns: 2,
	}
	client := &http.Client{
		Timeout:15*time.Second,
		Transport:transport,
	}
        request, err := http.NewRequest("{{.Method}}","{{.Url}}" , nil)
        var head[]string
        {{range .Headers}}
        head=strings.Split("{{ . }}",":")
        request.Header.Set(head[0], head[1])
        {{end}}
	resp,err:=client.Do(request)
	if err!=nil{
                return 
        }else{
             resp.Body.Close()
             return resp.StatusCode
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
	},nil
}
func (c *{{.Type}}Metrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.{{.Type}}Desc,
		prometheus.CounterValue,
		float64(StatusCode),
		"{{.LabelValues}}",
	)
	return nil
}
