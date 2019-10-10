.PHONY: build
NODEURL = $(PWD)/node_exporter

env:
	docker build -t monitor/centos7:v1 .
build:
	go build
	./monitor_yig
	cd $(NODEURL) && go build
	cd $(NODEURL) &&  ./node_exporter
run:
	docker run --name monitor_yig -v /root/monitor_yig/config:/root/monitor_yig/config  -p 9100:9100 -p 9999:9999 monitor/centos7:v1 /bin/bash -c 'make'

clean:
	docker rm -f  monitor_yig
log:
	journalctl -f -u docker.service
