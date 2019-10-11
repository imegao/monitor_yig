FROM centos:7.6.1810
WORKDIR /root
ENV GOROOT=/usr/local/go
ENV PATH=$GOROOT/bin:$PATH
RUN yum install -y make 
RUN yum install -y wget
RUN yum install -y git
RUN wget https://dl.google.com/go/go1.12.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.12.linux-amd64.tar.gz && rm -f go1.12.linux-amd64.tar.gz
RUN mkdir monitor_yig 
WORKDIR /root/monitor_yig

