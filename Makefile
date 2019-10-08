.PHONY: build
NODEURL = $(PWD)/node_exporter

build:
	go build
	./demo
	cd $(NODEURL) && go build
	cd $(NODEURL) && sudo ./node_exporter
