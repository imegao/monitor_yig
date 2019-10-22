package collector
import(
    "fmt"
    "os/exec"
    "strings"

    "github.com/prometheus/client_golang/prometheus"
)

type dnsmasq_01Metrics struct {
    dnsmasq_01Desc *prometheus.Desc
}

func (c *dnsmasq_01Metrics) process_function() (flag bool) {
    cmd:=exec.Command("systemctl","status","dnsmasq")
    out,err:=cmd.Output()
    if err !=nil{
        fmt.Println("systemctl dnsmasq err ",err)
    }
    flag=strings.Contains(string(out), "(running)")
    return flag
}

func init() {
	registerCollector("dnsmasq_01", defaultEnabled, Newdnsmasq_01Metrics)
}
func Newdnsmasq_01Metrics()  (Collector, error) {
    return &dnsmasq_01Metrics{
        dnsmasq_01Desc: prometheus.NewDesc(
            "process_dnsmasq_01_status",
            "process_dnsmasq_01_monitor",
            []string{"itemId","host"},
            nil,
         ),
    }, nil
}
func (c *dnsmasq_01Metrics) Update(ch chan<- prometheus.Metric) error{
    var value=0
    flag:=c.process_function()
    if flag==true{
        value=1
    } else{
        value=0
    }
    ch <- prometheus.MustNewConstMetric(
    	c.dnsmasq_01Desc,
        prometheus.CounterValue,
        float64(value),
        "dnsmasq_01",
        "192.168.2.128",
     )
     return nil
}



