package collector
import(
    "fmt"
    "os/exec"
    "strings"

    "github.com/prometheus/client_golang/prometheus"
)

type dnsmasqMetrics struct {
    dnsmasqDesc *prometheus.Desc
}

func (c *dnsmasqMetrics) process_function() (flag bool) {
    cmd:=exec.Command("systemctl","status","dnsmasq")
    out,err:=cmd.Output()
    if err !=nil{
        fmt.Println("systemctl dnsmasq err ",err)
    }
    flag=strings.Contains(string(out), "(running)")
    return flag
}

func init() {
	registerCollector("dnsmasq", defaultEnabled, NewdnsmasqMetrics)
}
func NewdnsmasqMetrics()  (Collector, error) {
    return &dnsmasqMetrics{
        dnsmasqDesc: prometheus.NewDesc(
            "process_dnsmasq_status",
            "process_dnsmasq_monitor",
            []string{"itemId","host"},
            nil,
         ),
    }, nil
}
func (c *dnsmasqMetrics) Update(ch chan<- prometheus.Metric) error{
    var value=0
    flag:=c.process_function()
    if flag==true{
        value=1
    } else{
        value=0
    }
    ch <- prometheus.MustNewConstMetric(
    	c.dnsmasqDesc,
        prometheus.CounterValue,
        float64(value),
        "dnsmasq",
        "192.168.2.128",
     )
     return nil
}



