package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type {{.Name}}Metrics struct {
    {{.Name}}Desc *prometheus.Desc
}

func (c *{{.Name}}Metrics) Http_function() (http_status int) {
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
	a:=strings.Split("{{.ReqHead}}","|")
        request, err := http.NewRequest("{{.ReqWay}}","{{.Url}}" , nil)
	for _,j:=range a{
            if j!=""{
	    	head:=strings.Split(j,":")
                request.Header.Set(head[0], head[1])
            }
         }


	resp,err:=client.Do(request)
	if err!=nil{
                return 
        }else{
             resp.Body.Close()
             return resp.StatusCode
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
	},nil
}
func (c *{{.Name}}Metrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.{{.Name}}Desc,
		prometheus.CounterValue,
		float64(StatusCode),
		"{{.LabelValues}}",
	)
	return nil
}
