# 对象存储系统自定义监控
监控内容主要分为5个部分：

* 机器硬件监控
* 服务进程监控
* 数据库监控
* HTTP监控
* TCP监控

监控服务主要基于prometheus、node-exporter、alertmanager、grafana这四种监控工具实现。

## 安装及使用
1、安装golang环境(版本1.11以上)

2、安装prometheus、alertmanager、grafana四种工具并启动。

3、下载monitor-yig进入config目录，配置config.yaml，执行make build编译，再执行，make run 运行。

## 配置文件说明
配置文件包括 prometheus配置文件、alertmanager配置文件、monitor\_yig配置文件，配置yaml文件要注意格式问题，大部分出错都是和格式有关，有时候直接复制粘贴也会引起不知名格式问题。

* 大小写敏感
* 配置文件冒号后要有空格
* 使用缩进表示层级关系
* 缩进时不能使用Tab键，只允许使用空格
* 缩进空格的数目不重要，只要相同层级的元素左侧对齐即可

### Prometheus配置说明

```YAML
#Prometheus全局配置项
global:
  scrape_interval:     15s # 设定抓取数据的周期，默认为1min
  evaluation_interval: 15s # 设定更新rules文件的周期，默认为1min
  scrape_timeout: 15s # 设定抓取数据的超时时间，默认为10s
  external_labels: # 额外的属性，会添加到拉取得数据并存到数据库中
   monitor: 'codelab_monitor'
#Alertmanager配置
alerting:
 alertmanagers:
 - static_configs:
   - targets: ["localhost:9093"] # 设定alertmanager和prometheus交互的接口，即alertmanager监听的ip地址和端口
     
#rule配置，首次读取默认加载，之后根据evaluation_interval设定的周期加载
rule_files:
 - "alertmanager_rules.yml"
 - "prometheus_rules.yml"

#scape配置
scrape_configs:
- job_name: 'prometheus' # job_name默认写入timeseries的labels中，可以用于查询使用
  scrape_interval: 15s # 抓取周期，默认采用global配置
  static_configs: # 静态配置
  - targets: ['localdns:9090'] # prometheus所要抓取数据的地址，即instance实例项

- job_name: 'example-random'
  static_configs:
  - targets: ['localhost:8080']
```
### Alertmanager配置说明

```YAML
# 全局配置项
global: 
  resolve_timeout: 5m #处理超时时间，默认为5min
  smtp_smarthost: 'smtp.sina.com:25' # 邮箱smtp服务器代理
  smtp_from: '******@sina.com' # 发送邮箱名称
  smtp_auth_username: '******@sina.com' # 邮箱名称
  smtp_auth_password: '******' #邮箱密码

# 定义路由树信息
route:
  group_by: ['alertname'] # 报警分组名称
  group_wait: 10s # 最初即第一次等待多久时间发送一组警报的通知
  group_interval: 10s # 在发送新警报前的等待时间
  repeat_interval: 1m # 发送重复警报的周期
  receiver: 'email' # 发送警报的接收者的名称，以下receivers name的名称

# 定义警报接收者信息
receivers:
  - name: 'email' # 警报
    email_configs: # 邮箱配置
    - to: '******@163.com'  # 接收警报的email配置

```
### Prometheus模块定义告警规则

```YAML
groups:
 - name: test-rules
   rules:
   - alert: InstanceDown # 告警名称
     expr: up == 0 # 告警的判定条件，参考Prometheus高级查询来设定
     for: 2m # 满足告警条件持续时间多久后，才会发送告警
     labels: #标签项
      team: node
     annotations: # 解析项，详细解释告警信息
      summary: "{{$labels.instance}}: has been down"
      description: "{{$labels.instance}}: job {{$labels.job}} has been down "
```
### Monitor_yig配置
配置文件配置内容为：通过模板需要生成监控项，以及监控项相关参数。

配置文件出错会导致编译不成功或生成监控文件出错，每次启动程序会根据配置文件自动重新生成相关监控文件如无法自动删除生成监控文件可以在node-expoeter/collector中手动删除。

运行参数 ./monitor-yig --p config.yaml 默认参数为config/config.yaml

* `targetPath:`通过监控模板生成监控文件的存放路径，一般存放在node-exporter的collector目录下，默认该项目目录下node-exporter。
* `databases、caches、https、processes、tcps:`通过5个参数来配置需要监控的内容。
* `~HostId：`该项区分不同的监控主机，不同监控主机的ID。
* `fileName:`通过模板生成监控文件的名称，必须唯一，不能重复，建议格式~_monitor.go。
* `Type:`为监控函数名称，必须唯一，不能重复，应为字母组合，不能为数字或特殊字符。
* `fqName:`为metrics采集数据的格式名称。
* `variableLabels:`为metrics采集数据的格式标签名称。
* `labelValues:`为metrics采集数据的格式标签值。
* `dataSourceName:`数据库登录信息(只在databases中配置)。
* `addr:`IP和端口（只在caches、tcps中配置）。
* `method:`http请求方法（只在https中配置）。
* `url:`域名配置（只在https中配置）。
* `Headers:`请求头及参数，使用数组的方式，请求头和请求参数使用‘:’分割。
* `password:`登录redis的密码(只在caches中配置)。

```YAML
#配置生成监控文件的路径
targetPath: ./node_exporter/collector/
#配置监控数据库内容
databases:#数组 不同主机
 - databaseHostId: 192.168.2.1   #该项区分不同监控主机，不同监控主机的ID。
   databaseNodes:  #数组 同一主机下配置不同的监控节点
   - fileName: tidb_monitor.go #通过模板生成监控文件的名称，必须唯一，不能重复，建议格式~_monitor.go。
     type: tidb #为监控函数名称，必须唯一，不能重复，应为字母组合，不能为数字或特殊字符
     dataSourceName: root:@tcp(192.168.2.128:4000)/yig?charset=utf8 #数据库登录信息
     fqName: tidb_status #为metrics采集数据的格式名称
     variableLabels: status #为metrics采集数据的格式标签名称
     labelValues: tidb #为metrics采集数据的格式标签值
#配置监控redis内容
caches:
  - cacheHostId: 192.168.2.1
    cacheNodes:
      - fileName: redis_monitor.go
        type: redis
        addr: 192.168.2.128:6379  #redis的IP
        password: hehehehe #登录密码
        fqName: redis_status
        variableLabels: status
        labelValues: redis
#配置监控HTTP请求内容
https:
  - httpHostId: 192.168.2.1
    httpNodes:
      - fileName: iam_api_monitor.go
        type: iam
        method: GET     #请求方式
        url: http://www.baidu.com  #请求URL
        headers: ["Accept-Language:zh-cn,en","dsafaf:fsvfds"]
        fqName: iam_status
        variableLabels: status
        labelValues: iam
      - fileName: postPay_api_monitor.go
        type: postPay
        method: GET
        url: http://www.baidu.com
        headers: ["Accept-Language:zh-cn,en","dsafaf:fsvfds"]
        fqName: postPay_status
        variableLabels: status
        labelValues: postPay
#配置监控systemctl服务
processes:
  - processHostId: 192.168.2.1
    processNodes:
      - fileName: process_monitor.go
        type: dnsmasq
        fqName: dnsmasq_status
        variableLabels: status
        labelValues: dnsmasq
#配置监控TCP连接
tcps:
  - tcpHostId: 192.168.2.1
    tcpNodes:
      - fileName: redistcp_monitor.go
        type: redistcp
        addr: 192.168.2.128:6379
        fqName: redistcp_status
        variableLabels: status
        labelValues: tcp
  - tcpHostId: 192.168.2.2
    tcpNodes:
      - fileName: tcp1_monitor.go
        type: tcp1
        addr: 192.168.2.128:9999
        fqName: tcp1_status
        variableLabels: status
        labelValues: tcp
```
## grafana模板
grafana监控面板模板文件为monitor_yig.json,浏览器打开grafana监页面导入。
## Docker运行
1、docker build -t monitor/centos7:v1 . 生成docker镜像(也可以使用make image)

2、docker run --name monitor\_yig -d -v /root/monitor\_yig/config:/root/monitor\_yig/config  -p 9100:9100 -p 9999:9999 monitor/centos7:v1 /bin/bash -c 'make build'(也可以使用make docker)

3、docker rm -f  monitor_yig 删除运行docker容器(也可以使用 make clean)

在docker环境中第一次编译运行需要加载golang包，如果在下载golang包过程中因网速问题卡住，需要docker restart 

