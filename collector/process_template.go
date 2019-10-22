package collector
import(
    "fmt"
    "os/exec"
    "strings"

    "github.com/prometheus/client_golang/prometheus"
)

type {{.ItemId}}Metrics struct {
    {{.ItemId}}Desc *prometheus.Desc
}

func (c *{{.ItemId}}Metrics) process_function() (flag bool) {
    cmd:=exec.Command("systemctl","status","{{.ItemId}}")
    out,err:=cmd.Output()
    if err !=nil{
        fmt.Println("systemctl {{.ItemId}} err ",err)
    }
    flag=strings.Contains(string(out), "(running)")
    return flag
}

func init() {
	registerCollector("{{.ItemId}}", defaultEnabled, New{{.ItemId}}Metrics)
}
func New{{.ItemId}}Metrics()  (Collector, error) {
    return &{{.ItemId}}Metrics{
        {{.ItemId}}Desc: prometheus.NewDesc(
            "process_{{.ItemId}}_status",
            "process_{{.ItemId}}_monitor",
            []string{"itemId","host"},
            nil,
         ),
    }, nil
}
func (c *{{.ItemId}}Metrics) Update(ch chan<- prometheus.Metric) error{
    var value=0
    flag:=c.process_function()
    if flag==true{
        value=1
    } else{
        value=0
    }
    ch <- prometheus.MustNewConstMetric(
    	c.{{.ItemId}}Desc,
        prometheus.CounterValue,
        float64(value),
        "{{.ItemId}}",
        "{{.Host}}",
     )
     return nil
}



