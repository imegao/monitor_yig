package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
var CONFIG CategoryConf
type CategoryConf struct{
	TargetPath string        `yaml:"targetPath"`
	Databases [] Database
	Caches [] Cache
	Https [] Http
	Processes [] Process
	Tcps[] Tcp
}
type Database struct {
	DatabaseId string   `yaml:"databaseId"`
	DatabaseNodes [] DatabaseNode   `yaml:"databaseNodes"`
}
type Cache struct {
	CacheId string    `yaml:"cacheId"`
	CacheNodes [] CacheNode `yaml:"cacheNodes"`
}

type Http struct {
	HttpId string      `yaml:"httpId"`
	HttpNodes [] HttpNode   `yaml:"httpNodes"`
}

type Tcp struct {
	TcpId string      `yaml:"tcpId"`
	TcpNodes [] TcpNode   `yaml:"tcpNodes"`
}
type Process struct {
	ProcessId string      `yaml:"processId"`
	ProcessNodes [] ProcessNode   `yaml:"processNodes"`
}
type DatabaseNode struct{
	FileName string   `yaml:"fileName"`
	Name string       `yaml:"name"`
	DataSourceName string   `yaml:"dataSourceName"`
	FqName string     `yaml:"fqName"`
	VariableLabels string `yaml:"variableLabels"`
	LabelValues string `yaml:"labelValues"`
}

type CacheNode struct{
	FileName string   `yaml:"fileName"`
        Name string       `yaml:"name"`
	Addr string       `yaml:"addr"`
	Password string   `yaml:"password"`
	FqName string     `yaml:"fqName"`
	VariableLabels string  `yaml:"variableLabels"`
	LabelValues string  `yaml:"labelValues"`
}

type HttpNode struct{
	FileName string   `yaml:"fileName"`
	Name string       `yaml:"name"`
	ReqWay string     `yaml:"reqWay"`
        Url string        `yaml:"url"`
        ReqHead string    `yaml:"reqHead"`
	FqName string     `yaml:"fqName"`
	VariableLabels string  `yaml:"variableLabels"`
	LabelValues string  `yaml:"labelValues"`
}

type TcpNode struct{
	FileName string   `yaml:"fileName"`
	Name string       `yaml:"name"`
	Addr string         `yaml:"addr"`
	FqName string     `yaml:"fqName"`
	VariableLabels string  `yaml:"variableLabels"`
	LabelValues string  `yaml:"labelValues"`
}
type ProcessNode struct{
	FileName string   `yaml:"fileName"`
	Name string       `yaml:"name"`
	FqName string     `yaml:"fqName"`
	VariableLabels string  `yaml:"variableLabels"`
	LabelValues string  `yaml:"labelValues"`
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
