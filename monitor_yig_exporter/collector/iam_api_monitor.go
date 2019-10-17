package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type iamMetrics struct {
    iamDesc *prometheus.Desc
}

func (c *iamMetrics) Http_function() (http_status int) {
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
	registerCollector("iam", defaultEnabled, NewiamMetrics)
}
func NewiamMetrics()  (Collector, error) {
	return &iamMetrics{
		iamDesc: prometheus.NewDesc(
		    "iam_status",
		    "iam_status_monitor",
			[]string{"status"},
			nil,
		),
	},nil
}
func (c *iamMetrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.iamDesc,
		prometheus.CounterValue,
		float64(StatusCode),
		"iam",
	)
	return nil
}
