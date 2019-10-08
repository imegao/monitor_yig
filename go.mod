module demo

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.37.4

	github.com/Sirupsen/logrus => github.com/sirupsen/logrus v1.4.1

	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190404164418-38d8ce5564a5
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190409044807-56b785ea58b2
	golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190327163128-167ebed0ec6d
	golang.org/x/net => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190405154228-4b34438f7a67
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190408220357-e5b8258f4918
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.3.2
	google.golang.org/appengine => github.com/golang/appengine v1.5.0

	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190404172233-64821d5d2107
	google.golang.org/grpc => github.com/grpc/grpc-go v1.19.1
)

require (
	github.com/beevik/ntp v0.2.0
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/ema/qdisc v0.0.0-20180104102928-b307c22d3ce7
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/godbus/dbus v0.0.0-20190402143921-271e53dc4968
	github.com/hodgesds/perf-utils v0.0.7
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/lufia/iostat v0.0.0-20170605150913-9f7362b77ad3
	github.com/mattn/go-xmlrpc v0.0.1
	github.com/mdlayher/genetlink v0.0.0-20181016160152-e97704c1b795 // indirect
	github.com/mdlayher/netlink v0.0.0-20181210160939-e069752bc835 // indirect
	github.com/mdlayher/wifi v0.0.0-20180727163819-efdf3f4195d9
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/prometheus/client_golang v1.0.0
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90
	github.com/prometheus/common v0.4.1
	github.com/prometheus/procfs v0.0.4
	github.com/siebenmann/go-kstat v0.0.0-20160321171754-d34789b79745
	github.com/sirupsen/logrus v1.4.2 // indirect
	github.com/soundcloud/go-runit v0.0.0-20150630195641-06ad41a06c4a
	go.uber.org/atomic v1.3.2 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v2 v2.2.1
)
