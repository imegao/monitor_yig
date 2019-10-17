package collector
import(
    "fmt"
    "os/exec"
    "strings"

    "github.com/prometheus/client_golang/prometheus"
)

type {{.Type}}Metrics struct {
    {{.Type}}Desc *prometheus.Desc
}

func (c *{{.Type}}Metrics) process_function() (flag bool) {
    cmd:=exec.Command("systemctl","status","{{.Type}}")
    out,err:=cmd.Output()
    if err !=nil{
        fmt.Println("systemctl {{.Type}} err ",err)
    }
    flag=strings.Contains(string(out), "(running)")
    return flag
}

func init() {
	registerCollector("{{.Type}}", defaultEnabled, New{{.Type}}Metrics)
}
func New{{.Type}}Metrics()  (Collector, error) {
    return &{{.Type}}Metrics{
        {{.Type}}Desc: prometheus.NewDesc(
            "{{.FqName}}",
            "{{.FqName}}_monitor",
            []string{"{{.VariableLabels}}"},
            nil,
         ),
    }, nil
}
func (c *{{.Type}}Metrics) Update(ch chan<- prometheus.Metric) error{
    var value=0
    flag:=c.process_function()
    if flag==true{
        value=1
    } else{
        value=0
    }
    ch <- prometheus.MustNewConstMetric(
    	c.{{.Type}}Desc,
        prometheus.CounterValue,
        float64(value),
        "{{.LabelValues}}",
     )
     return nil
}



