module github.com/prometheus/node_exporter

replace (
	go.uber.org/atomic => github.com/uber-go/atomic v1.4.0
	go.uber.org/multierr => github.com/uber-go/multierr v1.2.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190404164418-38d8ce5564a5
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190409044807-56b785ea58b2
	golang.org/x/image => github.com/golang/image v0.0.0-20190321063152-3fc05d484e9f
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190313153728-d0100b6bd8b3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190327163128-167ebed0ec6d
	golang.org/x/net => github.com/golang/net v0.0.0-20190404232315-eb5bcb51f2a3
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190902133755-9109b7679e13
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190408220357-e5b8258f4918
)

require (
	github.com/beevik/ntp v0.2.0
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e
	github.com/ema/qdisc v0.0.0-20190904071900-b82c76788043
	github.com/go-redis/redis v6.15.5+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/godbus/dbus v0.0.0-20190402143921-271e53dc4968
	github.com/hodgesds/perf-utils v0.0.7
	github.com/lufia/iostat v0.0.0-20170605150913-9f7362b77ad3
	github.com/mattn/go-xmlrpc v0.0.1
	github.com/mdlayher/wifi v0.0.0-20190303161829-b1436901ddee
	github.com/prometheus/client_golang v1.11.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.26.0
	github.com/prometheus/procfs v0.6.0
	github.com/siebenmann/go-kstat v0.0.0-20160321171754-d34789b79745
	github.com/soundcloud/go-runit v0.0.0-20150630195641-06ad41a06c4a
	golang.org/x/sys v0.4.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mdlayher/genetlink v0.0.0-20190828143517-e35f2bf499b9 // indirect
	github.com/mdlayher/netlink v0.0.0-20190828143259-340058475d09 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.26.0 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/stretchr/testify v1.5.1 // indirect
	go.uber.org/atomic v0.0.0-00010101000000-000000000000 // indirect
	go.uber.org/multierr v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
