package collector
import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
        "time"
        "net"
        "strings"
)
type postPayMetrics struct {
    postPayDesc *prometheus.Desc
}

func (c *postPayMetrics) Http_function() (http_status int) {
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
	a:=strings.Split("Accept-Language:zh-cn,zh|Accept-Encoding:gzip,deflate|fhdsvhodv:fdsvfsv|Cookie:JSESSIONID=369766FDF6220F7803433C0B2DE36D98","|")
        request, err := http.NewRequest("HEAD","http://192.168.2.128:7777/hello?user=admin&pass=777" , nil)
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
	registerCollector("postPay", defaultEnabled, NewpostPayMetrics)
}
func NewpostPayMetrics()  (Collector, error) {
	return &postPayMetrics{
		postPayDesc: prometheus.NewDesc(
		    "postPay_status",
		    "postPay_status_monitor",
			[]string{"status"},
			nil,
		),
	},nil
}
func (c *postPayMetrics) Update(ch chan<- prometheus.Metric) error{
	StatusCode:= c.Http_function()
	ch <- prometheus.MustNewConstMetric(
		c.postPayDesc,
		prometheus.CounterValue,
		float64(StatusCode),
		"postPay",
	)
	return nil
}
