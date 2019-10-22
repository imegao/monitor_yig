package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type Iam_0001Metrics struct {
    Iam_0001Desc *prometheus.Desc
}

func (c *Iam_0001Metrics) Http_function() (http_status int) {
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
        request, err := http.NewRequest("GET","http://192.168.2.128:8888/hello?user=admin&pass=888" , nil)
        var head[]string
        
        head=strings.Split("Accept-Language:zh-cn,en",":")
        request.Header.Set(head[0], head[1])
        
        head=strings.Split("dsafaf:fsvfds",":")
        request.Header.Set(head[0], head[1])
        
	resp,err:=client.Do(request)
	if err!=nil{
                return 
        }else{
             resp.Body.Close()
             return resp.StatusCode
        }
}

func init() {
	registerCollector("Iam_0001", defaultEnabled, NewIam_0001Metrics)
}
func NewIam_0001Metrics()  (Collector, error) {
	return &Iam_0001Metrics{
		Iam_0001Desc: prometheus.NewDesc(
		    "http_Iam_0001_status",
		    "http_Iam_0001_monitor",
		    []string{"itemId","host"},
		    nil,
		),
	},nil
}
func (c *Iam_0001Metrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.Iam_0001Desc,
		prometheus.CounterValue,
		float64(StatusCode),
		"Iam_0001",
                "192.168.2.128",
	)
	return nil
}
