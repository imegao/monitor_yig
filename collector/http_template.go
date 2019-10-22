package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type {{.ItemId}}Metrics struct {
    {{.ItemId}}Desc *prometheus.Desc
}

func (c *{{.ItemId}}Metrics) Http_function() (http_status int) {
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
	registerCollector("{{.ItemId}}", defaultEnabled, New{{.ItemId}}Metrics)
}
func New{{.ItemId}}Metrics()  (Collector, error) {
	return &{{.ItemId}}Metrics{
		{{.ItemId}}Desc: prometheus.NewDesc(
		    "http_{{.ItemId}}_status",
		    "http_{{.ItemId}}_monitor",
		    []string{"itemId","host"},
		    nil,
		),
	},nil
}
func (c *{{.ItemId}}Metrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.{{.ItemId}}Desc,
		prometheus.CounterValue,
		float64(StatusCode),
		"{{.ItemId}}",
                "{{.Host}}",
	)
	return nil
}
