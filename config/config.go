package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
var CONFIG CategoryConf
type CategoryConf struct{
	TargetPath string        `yaml:"targetPath"`
	Host string              `yaml:"host"`
        Mysql   [] MysqlNode
	Redis   [] RedisNode
	Http    [] HttpNode
	Process [] ProcessNode
	Tcp     [] TcpNode
}
type MysqlNode struct{
	ItemId         string    `yaml:"itemId"`
        DataSourceName string    `yaml:"dataSourceName"`
        Host           string
}

type RedisNode struct{
        ItemId         string    `yaml:"itemId"`
	Addr           string    `yaml:"addr"`
	Password       string    `yaml:"password"`
        Host           string
}

type HttpNode struct{
        ItemId         string    `yaml:"itemId"`
	Method         string    `yaml:"method"`
        Url            string    `yaml:"url"`
        Headers        []string  `yaml:"headers"`
        Host           string
}

type TcpNode struct{
        ItemId         string    `yaml:"itemId"`
	Addr           string    `yaml:"addr"`
        Host           string
}
type ProcessNode struct{
        ItemId         string    `yaml:"itemId"`
        Host           string
}
func ReadConfig(p string){
	data, err := ioutil.ReadFile(p)
	if err!=nil{
		fmt.Println(err)
	}

	//fmt.Println(string(data))
	//var c CategoryConf
	err=yaml.Unmarshal(data, &CONFIG)
	if err!=nil{
		fmt.Println("y",err)
	}
	fmt.Println("ConfigData", CONFIG)
	//d, _ := yaml.Marshal(&t)
      
}
