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
	Httpservers [] Http
	Processes [] Process
	Tcps[] Tcp
}
type Database struct {
	DatabaseHostId string   `yaml:"databaseHostId"`
	DatabaseNodes [] DatabaseNode   `yaml:"databaseNodes"`
}
type Cache struct {
	CacheHostId string    `yaml:"cacheHostId"`
	CacheNodes [] CacheNode `yaml:"cacheNodes"`
}

type Http struct {
	HttpHostId string      `yaml:"httpHostId"`
	HttpNodes [] HttpNode   `yaml:"httpNodes"`
}

type Tcp struct {
	TcpHostId string      `yaml:"tcpHostId"`
	TcpNodes [] TcpNode   `yaml:"tcpNodes"`
}
type Process struct {
	ProcessHostId string      `yaml:"processHostId"`
	ProcessNodes [] ProcessNode   `yaml:"processNodes"`
}
type DatabaseNode struct{
	FileName string   `yaml:"fileName"`
	Type string       `yaml:"type"`
	DataSourceName string   `yaml:"dataSourceName"`
	FqName string     `yaml:"fqName"`
	VariableLabels string `yaml:"variableLabels"`
	LabelValues string `yaml:"labelValues"`
}

type CacheNode struct{
	FileName string   `yaml:"fileName"`
        Type string       `yaml:"type"`
	Addr string       `yaml:"addr"`
	Password string   `yaml:"password"`
	FqName string     `yaml:"fqName"`
	VariableLabels string  `yaml:"variableLabels"`
	LabelValues string  `yaml:"labelValues"`
}

type HttpNode struct{
	FileName string   `yaml:"fileName"`
	Type string       `yaml:"type"`
	Method string     `yaml:"method"`
        Url string        `yaml:"url"`
        Headers []string    `yaml:"headers"`
	FqName string     `yaml:"fqName"`
	VariableLabels string  `yaml:"variableLabels"`
	LabelValues string  `yaml:"labelValues"`
}

type TcpNode struct{
	FileName string   `yaml:"fileName"`
	Type string       `yaml:"type"`
	Addr string         `yaml:"addr"`
	FqName string     `yaml:"fqName"`
	VariableLabels string  `yaml:"variableLabels"`
	LabelValues string  `yaml:"labelValues"`
}
type ProcessNode struct{
	FileName string   `yaml:"fileName"`
	Type string       `yaml:"type"`
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
