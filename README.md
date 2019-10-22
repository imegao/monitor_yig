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

3、下载monitor-yig进入config目录，配置config.yaml，执行make build编译，再执行make run 运行,make stop停止（使用make stop 注意避免将本机其他node_exporter关闭）。

4、在Makefile中配置exporter的端口，默认端口是9100
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
* `host:`当前主机的IP
* `mysql、redis、http、process、tcp:`通过5个参数来配置需要监控的类型。
* `itemId:`为监控项的ID,必须唯一，格式为（字符串+四位数字）
* `dataSourceName:`数据库登录信息(只在databases中配置)。
* `addr:`IP和端口（只在caches、tcps中配置）。
* `method:`http请求方法（只在https中配置）。
* `url:`域名配置（只在https中配置）。
* `Headers:`请求头及参数，使用数组的方式，请求头和请求参数使用‘:’分割。
* `password:`登录redis的密码(只在caches中配置)。

```YAML
#配置生成监控文件的路径
targetPath: ./monitor_yig_exporter/collector/
host: 192.168.2.128  #监控主机的IP
mysql:
  - itemId: tidb_4000  监控项的ID,必须唯一，格式：字符串_四位数字
    dataSourceName: root:@tcp(192.168.2.128:4000)/yig?charset=utf8
redis:
  - itemId: redis_6379
    addr: 192.168.2.128:6379
    password: hehehehe
http:
  - itemId: Iam_0001
    method: GET
    url: http://192.168.2.128:8888/hello?user=admin&pass=888
    headers: ["Accept-Language:zh-cn,en","dsafaf:fsvfds"]
  - itemId: PostPay_0001
    method: HEAD
    url: http://192.168.2.128:7777/hello?user=admin&pass=777
    headers: ["Accept-Language:zh-cn,zh","Accept-Encoding:gzip,deflate","fhdsvhodv:fdsvfsv","Cookie:JSESSIONID=369766FDF6220F7803433C0B2DE36D98"]
process:
  - itemId: dnsmasq  #监控systemctl服务的名称 不同于其他ID
tcp:
  - itemId: yig_8080
    addr: 192.168.2.128:8080
  - itemId: yig_9000
    addr: 192.168.2.128:9000
```
## grafana模板
grafana监控面板模板文件为monitor_yig.json,浏览器打开grafana监页面导入。
## Docker运行
1、docker build -t monitor/centos7:v1 . 生成docker镜像(也可以使用make image)

2、docker run --name monitor\_yig -d -v /root/monitor\_yig/config:/root/monitor\_yig/config  -p 9100:9100 -p 9999:9999 monitor/centos7:v1 /bin/bash -c 'make build'(也可以使用make docker)

3、docker rm -f  monitor_yig 删除运行docker容器(也可以使用 make clean)

在docker环境中第一次编译运行需要加载golang包，如果在下载golang包过程中因网速问题卡住，需要docker restart 

