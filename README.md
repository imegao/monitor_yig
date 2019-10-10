# 对象存储系统自定义监控
监控内容主要分为5个部分：

* 机器硬件监控
* 服务进程监控
* 数据库监控
* HTTP监控
* TCP监控

监控服务主要基于prometheus、node-exporter、alertmanager、grafana这四种监控工具实现。

##安装及使用
1、安装golang环境

2、安装prometheus、alertmanager、grafana四种工具并启动。

3、下载monitor-yig进入config目录，配置config.yaml，执行make build。

## 配置文件说明
配置文件配置内容为：通过模板需要生成监控项，以及监控项相关参数。

配置文件出错会导致编译不成功或生成监控文件出错，每次启动程序会根据配置文件自动重新生成相关监控文件如无法自动删除生成监控文件可以在node-expoeter/collector中手动删除。

运行参数 ./monitor-yig --p ~.yaml 默认参数为config/config.yaml

* `targetPath:`通过监控模板生成监控文件的存放路径，一般存放在node-exporter的collector目录下，默认该项目目录下node-exporter。
* `databases、caches、https、processes、tcps:`通过5个参数来配置需要监控的内容。
* `~Id：`该项区分同一监控内容下不同的监控项目，不同监控项目的ID必须为唯一数字，不能重复。
* `fileName:`通过模板生成监控文件的名称，必须唯一，不能重复，建议格式~_monitor.go。
* `name:`为监控函数名称，必须唯一，不能重复，应为字母组合，不能为数字或特殊字符。
* `fqName:`为metrics采集数据的格式名称。
* `variableLabels:`为metrics采集数据的格式标签名称。
* `labelValues:`为metrics采集数据的格式标签值。
* `dataSourceName:`数据库登录信息(只在databases中配置)。
* `addr:`IP和端口（只在caches、tcps中配置）。
* `reqWay:`http请求方法（只在https中配置）。
* `url:`域名配置（只在https中配置）。
* `reqHead:`请求头及参数，请求头和请求参数使用‘:’分割，两条请求头参间使用‘|’分割（只在https中配置）。
* `password:`登录redis的密码(只在caches中配置)。
配置文件例子在文档最后
## grafana模板
grafana监控面板模板文件为monitor_yig.json,浏览器打开grafana监页面导入。
## Docker运行
1、docker build -t monitor/centos7:v1 . 生成docker镜像(也可以使用make image)

2、docker run --name monitor\_yig -d -v /root/monitor\_yig/config:/root/monitor\_yig/config  -p 9100:9100 -p 9999:9999 monitor/centos7:v1 /bin/bash -c 'make build'(也可以使用make run)

3、docker rm -f  monitor_yig 删除运行docker容器(也可以使用 make clean)

在docker环境中第一次编译运行需要加载golang包，如果在下载golang包过程中因网速问题卡住，需要docker restart 

## config.yaml
```YAML
targetPath: ./node_exporter/collector/
databases:
 - databaseId: 1
   databaseNodes:
   - fileName: tidb_monitor.go
     name: tidb
     dataSourceName: root:@tcp(192.168.2.128:4000)/yig?charset=utf8
     fqName: tidb_status
     variableLabels: status
     labelValues: tidb
caches:
  - cacheId: 1
    cacheNodes:
      - fileName: redis_monitor.go
        name: redis
        addr: 192.168.2.128:6379
        password: hehehehe
        fqName: redis_status
        variableLabels: status
        labelValues: redis
https:
  - httpId: 1
    httpNodes:
      - fileName: iam_api_monitor.go
        name: iam
        reqWay: GET
        url: http://www.baidu.com
        reqHead: Accept:text/html,application/xhtml+xml,application/xml\Accept-Encoding:gzip, deflate\
        fqName: iam_status
        variableLabels: status
        labelValues: iam
  - httpId: 2
    httpNodes:
      - fileName: postPay_api_monitor.go
        name: postPay
        reqWay: GET
        url: http://www.baidu.com
        reqHead: Accept-Language:zh-cn,zh;
        fqName: postPay_status
        variableLabels: status
        labelValues: postPay
processes:
  - processId: 1
    processNodes:
      - fileName: process_monitor.go
        name: dnsmasq
        fqName: dnsmasq_status
        variableLabels: status
        labelValues: dnsmasq
tcps:
  - tcpId: 1
    tcpNodes:
      - fileName: redistcp_monitor.go
        name: redistcp
        addr: 192.168.2.128:6379
        fqName: redistcp_status
        variableLabels: status
        labelValues: tcp
  - tcpId: 2
    tcpNodes:
      - fileName: tcp1_monitor.go
        name: tcp1
        addr: 192.168.2.128:9999
        fqName: tcp1_status
        variableLabels: status
        labelValues: tcp
```
