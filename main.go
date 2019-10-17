package main
import (
	"Generate_MonitorFiles/config"
	"flag"
	"fmt"
	"os"
	"text/template"
)

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
	for _,host:=range config.CONFIG.Databases{
		fmt.Println("DatabaseHostId: ",host.DatabaseHostId)
		for _,node:=range host.DatabaseNodes{
                    generateDb(&node)
                }
	}
	for _,host:=range config.CONFIG.Caches{
		fmt.Println("CacheHostId: ",host.CacheHostId)
		for _,node:=range host.CacheNodes{
                    generateCache(&node)
	        }
        }
	for _,host:=range config.CONFIG.Tcps{
		fmt.Println("TcpHostId: ",host.TcpHostId)
                for _,node:=range host.TcpNodes{
		    generateTcp(&node)
                }
	}
	for _,host:=range config.CONFIG.Httpservers{
		fmt.Println("HttpHostId: ",host.HttpHostId)
		for _,node:=range host.HttpNodes{
                    generateHttp(&node)
                }
	}
	for _,host:=range config.CONFIG.Processes{
		fmt.Println("ProcessHostId: ",host.ProcessHostId)
                for _,node:=range host.ProcessNodes{
		    generateProcess(&node)
                }
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
	}
}
func generateHttp(c *config.HttpNode){
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
	} else {
               fmt.Printf("generate %s success\n",c.FileName)
               }
	httpTemplate.Execute(f,c)
}
func generateCache(c *config.CacheNode){
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
	} else {
               fmt.Printf("generate %s success\n",c.FileName)
               }
	cacheTemplate.Execute(f,c)
}
func generateDb(c *config.DatabaseNode){
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
	} else {
               fmt.Printf("generate %s success\n",c.FileName)
               }
	dbTemplate.Execute(f,c)

}
func generateProcess(c *config.ProcessNode){
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
	} else {
               fmt.Printf("generate %s success\n",c.FileName)
               }
	processTemplate.Execute(f,c)

}
func generateTcp(c *config.TcpNode){
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
	} else {
               fmt.Printf("generate %s success\n",c.FileName)
               }
	tcpTemplate.Execute(f,c)

}
