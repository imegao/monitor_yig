package main
import (
	"demo/config"
	"flag"
	"fmt"
	"os"
	"text/template"
)

type Database struct {
	FileName string
	Name string
	DataSourceName  string
	FqName string
	VariableLabels string
	LabelValues string

}

type Cache struct {
	FileName string
	Name string
	Addr  string
	Password string
	FqName string
	VariableLabels string
	LabelValues string

}
type Tcp struct {
	FileName string
	Name string
	Addr  string
	FqName string
	VariableLabels string
	LabelValues string
}
type Http struct {
	FileName string
	Name string
	ReqWay string
	Url  string
	ReqHead string
	FqName string
	VariableLabels string
	LabelValues string
}
type Process struct {
	FileName string
	Name string
	FqName string
	VariableLabels string
	LabelValues string
}


func main() {
	var ConfigPath string
	flag.StringVar(&ConfigPath,"p","config/config.yaml","Path of configuration file")
	flag.Parse()
	fmt.Println(ConfigPath)
	config.ReadConfig(ConfigPath)
	_,err:=os.Stat(config.CONFIG.TargetPath)
	if err!=nil{
		fmt.Println(err)
	}
	for _,j:=range config.CONFIG.Databases{
		fmt.Println("DatabaseId: ",j.DatabaseId)
		test:=Database{j.DatabaseNodes[0].FileName,j.DatabaseNodes[0].Name,j.DatabaseNodes[0].DataSourceName,j.DatabaseNodes[0].FqName,j.DatabaseNodes[0].VariableLabels,j.DatabaseNodes[0].LabelValues}
		generateDb(&test)
	}
	for _,j:=range config.CONFIG.Caches{
		fmt.Println("CacheId: ",j.CacheId)
		test1:=Cache{j.CacheNodes[0].FileName,j.CacheNodes[0].Name,j.CacheNodes[0].Addr,j.CacheNodes[0].Password,j.CacheNodes[0].FqName,j.CacheNodes[0].VariableLabels,j.CacheNodes[0].LabelValues}
		generateCache(&test1)
	}
	for _,j:=range config.CONFIG.Tcps{
		fmt.Println("TcpId: ",j.TcpId)
		test:=Tcp{j.TcpNodes[0].FileName,j.TcpNodes[0].Name,j.TcpNodes[0].Addr,j.TcpNodes[0].FqName,j.TcpNodes[0].VariableLabels,j.TcpNodes[0].LabelValues}
		generateTcp(&test)
	}
	for _,j:=range config.CONFIG.Https{
		fmt.Println("HttpId: ",j.HttpId)
		test:=Http{j.HttpNodes[0].FileName,j.HttpNodes[0].Name,j.HttpNodes[0].ReqWay,j.HttpNodes[0].Url,j.HttpNodes[0].ReqHead,j.HttpNodes[0].FqName,j.HttpNodes[0].VariableLabels,j.HttpNodes[0].LabelValues}
		generateHttp(&test)
	}
	for _,j:=range config.CONFIG.Processes{
		fmt.Println("ProcessId: ",j.ProcessId)
		test:=Process{j.ProcessNodes[0].FileName,j.ProcessNodes[0].Name,j.ProcessNodes[0].FqName,j.ProcessNodes[0].VariableLabels,j.ProcessNodes[0].LabelValues}
		generateProcess(&test)
	}

}

func generateFile(filename string)(){
	_,err:=os.Stat(filename)

	if err==nil {
		err := os.Remove(filename)
		if err != nil {
			//如果删除失败则输出 file remove Error!
			fmt.Printf("%s remove Error!\n",filename)
			//输出错误详细信息
			fmt.Printf("%s", err)
		} else {
			//如果删除成功则输出 file remove OK!
			fmt.Printf("%s remove OK!\n",filename)
		}
	} else{
		fmt.Println(err)
		return
	}
}
func generateHttp(c *Http){
    var err error
	var httpTemplate *template.Template
	httpTemplate, err = template.ParseFiles("./collector/http_template.go")
	if err != nil {
		fmt.Println("HTTP parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+c.FileName)
	f,err:=os.OpenFile(config.CONFIG.TargetPath+c.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("HTTP open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	}

	httpTemplate.Execute(f,c)
}
func generateCache(c *Cache){
	var err error
	var cacheTemplate *template.Template
	cacheTemplate, err = template.ParseFiles("./collector/cache_template.go")
	if err != nil {
		fmt.Println("cache_template parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+c.FileName)
	f,err:=os.OpenFile(config.CONFIG.TargetPath+c.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("cache_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	}

	cacheTemplate.Execute(f,c)
}
func generateDb(c *Database){
	var err error
	var dbTemplate *template.Template
	dbTemplate, err = template.ParseFiles("./collector/db_template.go")
	if err != nil {
		fmt.Println("db_template parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+c.FileName)
	f,err:=os.OpenFile(config.CONFIG.TargetPath+c.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("db_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	}

	dbTemplate.Execute(f,c)

}
func generateProcess(c *Process){
	var err error
	var processTemplate *template.Template
	processTemplate, err = template.ParseFiles("./collector/process_template.go")
	if err != nil {
		fmt.Println("process_template parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+c.FileName)
	f,err:=os.OpenFile(config.CONFIG.TargetPath+c.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("process_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	}

	processTemplate.Execute(f,c)

}
func generateTcp(c *Tcp){
	var err error
	var tcpTemplate *template.Template
	tcpTemplate, err = template.ParseFiles("./collector/tcp_template.go")
	if err != nil {
		fmt.Println("tcp_template parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+c.FileName)
	f,err:=os.OpenFile(config.CONFIG.TargetPath+c.FileName,os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("tcp_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	}

	tcpTemplate.Execute(f,c)

}
