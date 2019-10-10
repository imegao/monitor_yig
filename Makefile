.PHONY: build
NODEURL = $(PWD)/node_exporter

image:
	docker build -t monitor/centos7:v1 .
build:
	go build
	./monitor_yig
	cd $(NODEURL) && go build
	cd $(NODEURL) &&  ./node_exporter
run:
	docker run --name monitor_yig -d -v /root/monitor_yig/config:/root/monitor_yig/config  -p 9100:9100 -p 9999:9999 monitor/centos7:v1 /bin/bash -c 'make build'

clean:
	docker rm -f  monitor_yig
log:
	journalctl -f -u docker.service
