package collector
import(
    "fmt"
    "os/exec"
    "strings"

    "github.com/prometheus/client_golang/prometheus"
)

type {{.Name}}Metrics struct {
    {{.Name}}Desc *prometheus.Desc
}

func (c *{{.Name}}Metrics) process_function() (flag bool) {
    cmd:=exec.Command("systemctl","status","{{.Name}}")
    out,err:=cmd.Output()
    if err !=nil{
        fmt.Println(err)
    }
    flag=strings.Contains(string(out), "(running)")
    return flag
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
    }, nil
}
func (c *{{.Name}}Metrics) Update(ch chan<- prometheus.Metric) error{
    var value=0
    flag:=c.process_function()
    if flag==true{
        value=1
    } else{
        value=0
    }
    ch <- prometheus.MustNewConstMetric(
    	c.{{.Name}}Desc,
        prometheus.CounterValue,
        float64(value),
        "{{.LabelValues}}",
     )
     return nil
}



