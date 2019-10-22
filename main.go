package main
import (
	"Generate_MonitorFiles/config"
	"flag"
	"fmt"
	"os"
	"text/template"
        "regexp"
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
	for _,node:=range config.CONFIG.Mysql{
                generateMysql(&node)
        }
        for _,node:=range config.CONFIG.Redis{
                generateRedis(&node)
        }
        for _,node:=range config.CONFIG.Tcp{
                generateTcp(&node)
        }
        for _,node:=range config.CONFIG.Http{
                generateHttp(&node)
        }
        for _,node:=range config.CONFIG.Process{
                generateProcess(&node)
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
        ok,_:=regexp.MatchString("^[a-zA-Z]+[_][0-9]{4}$",c.ItemId)
        if ok!=true{
                fmt.Println("HTTP itemId Configuration error")
                os.Exit(1)
        }
	httpTemplate, err = template.ParseFiles("./collector/http_template.go")
	if err != nil {
		fmt.Println("HTTP parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+"http_"+c.ItemId+"_monitor.go")
	f,err:=os.OpenFile(config.CONFIG.TargetPath+"http_"+c.ItemId+"_monitor.go",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("HTTP open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
               fmt.Printf("generate %s success\n","http_"+c.ItemId+"_monitor.go")
               }
        c.Host=config.CONFIG.Host
	httpTemplate.Execute(f,c)
}
func generateRedis(c *config.RedisNode){
	var err error
	var cacheTemplate *template.Template
        ok,_:=regexp.MatchString("^[a-zA-Z]+[_][0-9]{4}$",c.ItemId)
        if ok!=true{
                fmt.Println("Redis itemId Configuration error")
                os.Exit(1)
        }
	cacheTemplate, err = template.ParseFiles("./collector/redis_template.go")
	if err != nil {
		fmt.Println("cache_template parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+"redis_"+c.ItemId+"_monitor.go")
	f,err:=os.OpenFile(config.CONFIG.TargetPath+"redis_"+c.ItemId+"_monitor.go",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("redis_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
               fmt.Printf("generate %s success\n","redis_"+c.ItemId+"_monitor.go")
               }
        c.Host=config.CONFIG.Host
	cacheTemplate.Execute(f,c)
}
func generateMysql(c *config.MysqlNode){
	var err error
	var dbTemplate *template.Template
        ok,_:=regexp.MatchString("^[a-zA-Z]+[_][0-9]{4}$",c.ItemId)
        if ok!=true{
                fmt.Println("Mysql itemId Configuration error")
                os.Exit(1)
        }  
	dbTemplate, err = template.ParseFiles("./collector/mysql_template.go")
	if err != nil {
		fmt.Println("mysql_template parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+"mysql_"+c.ItemId+"_monitor.go")
	f,err:=os.OpenFile(config.CONFIG.TargetPath+"mysql_"+c.ItemId+"_monitor.go",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("mysql_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
               fmt.Printf("generate %s success\n","mysql_"+c.ItemId+"_monitor.go")
               }
        c.Host=config.CONFIG.Host
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
	generateFile(config.CONFIG.TargetPath+"process_"+c.ItemId+"_monitor.go")
	f,err:=os.OpenFile(config.CONFIG.TargetPath+"process_"+c.ItemId+"_monitor.go",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("process_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
               fmt.Printf("generate %s success\n","process_"+c.ItemId+"_monitor.go")
               }
        c.Host=config.CONFIG.Host
	processTemplate.Execute(f,c)
}
func generateTcp(c *config.TcpNode){
	var err error
	var tcpTemplate *template.Template
        ok,_:=regexp.MatchString("^[a-zA-Z]+[_][0-9]{4}$",c.ItemId)
        if ok!=true{
                fmt.Println("TCP itemId Configuration error")
                os.Exit(1)
        }  
	tcpTemplate, err = template.ParseFiles("./collector/tcp_template.go")
	if err != nil {
		fmt.Println("tcp_template parse file err:", err)
		return
	}
	generateFile(config.CONFIG.TargetPath+"tcp_"+c.ItemId+"_monitor.go")
	f,err:=os.OpenFile(config.CONFIG.TargetPath+"tcp_"+c.ItemId+"_monitor.go",os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer f.Close()
	if err != nil {
		fmt.Println("tcp_template file open Error!")
		//输出错误详细信息
		fmt.Printf("%s", err)
	} else {
               fmt.Printf("generate %s success\n","tcp_"+c.ItemId+"_monitor.go")
               }
        c.Host=config.CONFIG.Host
	tcpTemplate.Execute(f,c)

}
