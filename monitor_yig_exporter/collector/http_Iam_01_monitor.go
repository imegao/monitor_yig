package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type Iam_01Metrics struct {
    Iam_01Desc *prometheus.Desc
}

func (c *Iam_01Metrics) Http_function() (http_status int) {
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
	registerCollector("Iam_01", defaultEnabled, NewIam_01Metrics)
}
func NewIam_01Metrics()  (Collector, error) {
	return &Iam_01Metrics{
		Iam_01Desc: prometheus.NewDesc(
		    "http_Iam_01_status",
		    "http_Iam_01_monitor",
		    []string{"itemId","host"},
		    nil,
		),
	},nil
}
func (c *Iam_01Metrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.Iam_01Desc,
		prometheus.CounterValue,
		float64(StatusCode),
		"Iam_01",
                "192.168.2.128",
	)
	return nil
}
