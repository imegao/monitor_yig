.PHONY: build
NODEURL = $(PWD)/monitor_yig_exporter
CURL = $(NODEURL)/collector
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
image:
	docker build -t monitor/centos7:v1 .
build:  
	cd $(CURL) && rm -rf *monitor.go
	go build && ./Generate_MonitorFiles
	cd $(NODEURL) && go build
run:
	cd $(NODEURL) && nohup ./node_exporter --web.listen-address=":9100" &
stop:
	ps -ef |grep node_exporter |grep -v grep|cut -c 9-16|xargs kill -9
docker:
	docker run --name monitor_yig  --privileged=true -d -v /root/monitor_yig/node_exporter:/root/monitor_yig/ -p 9100:9100 -p 9999:9999 monitor/centos7:v1 /bin/bash -c './node_exporter'

clean:
	docker rm -f  monitor_yig
	cd $(CURL) && rm -rf *monitor.go
log:
	journalctl -f -u docker.service
