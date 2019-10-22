package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type PostPay_0001Metrics struct {
    PostPay_0001Desc *prometheus.Desc
}

func (c *PostPay_0001Metrics) Http_function() (http_status int) {
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
        request, err := http.NewRequest("HEAD","http://192.168.2.128:7777/hello?user=admin&pass=777" , nil)
        var head[]string
        
        head=strings.Split("Accept-Language:zh-cn,zh",":")
        request.Header.Set(head[0], head[1])
        
        head=strings.Split("Accept-Encoding:gzip,deflate",":")
        request.Header.Set(head[0], head[1])
        
        head=strings.Split("fhdsvhodv:fdsvfsv",":")
        request.Header.Set(head[0], head[1])
        
        head=strings.Split("Cookie:JSESSIONID=369766FDF6220F7803433C0B2DE36D98",":")
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
	registerCollector("PostPay_0001", defaultEnabled, NewPostPay_0001Metrics)
}
func NewPostPay_0001Metrics()  (Collector, error) {
	return &PostPay_0001Metrics{
		PostPay_0001Desc: prometheus.NewDesc(
		    "http_PostPay_0001_status",
		    "http_PostPay_0001_monitor",
		    []string{"itemId","host"},
		    nil,
		),
	},nil
}
func (c *PostPay_0001Metrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.PostPay_0001Desc,
		prometheus.CounterValue,
		float64(StatusCode),
		"PostPay_0001",
                "192.168.2.128",
	)
	return nil
}
