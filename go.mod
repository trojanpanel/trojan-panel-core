module trojan-panel-core

go 1.18

require (
	github.com/antonfisher/nested-logrus-formatter v1.3.1
	github.com/caddyserver/certmagic v0.16.1
	github.com/didi/gendry v1.7.0
	github.com/fsnotify/fsnotify v1.5.4
	github.com/gin-gonic/gin v1.8.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lucas-clemente/quic-go v0.28.0
	github.com/oschwald/geoip2-golang v1.7.0
	github.com/p4gefau1t/trojan-go v0.10.6
	github.com/prometheus/client_golang v1.12.2
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.12.0
	github.com/tobyxdd/hysteria v1.1.0
	github.com/xtls/xray-core v1.5.9
	github.com/yosuke-furukawa/json5 v0.1.1
	google.golang.org/grpc v1.48.0
	gopkg.in/ini.v1 v1.66.6
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/LiamHaworth/go-tproxy v0.0.0-20190726054950-ef7efd7f24ed // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/dgryski/go-metro v0.0.0-20211217172704-adc40b04c140 // indirect
	github.com/elazarl/goproxy v0.0.0-20220115173737-adb46da277ac // indirect
	github.com/elazarl/goproxy/ext v0.0.0-20220115173737-adb46da277ac // indirect
	github.com/eycorsican/go-tun2socks v1.16.11 // indirect
	github.com/francoispqt/gojay v1.2.13 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/gopacket v1.1.19 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.1.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/libdns/libdns v0.2.1 // indirect
	github.com/lunixbochs/struc v0.0.0-20200707160740-784aaebc1d40 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/marten-seemann/qtls-go1-16 v0.1.5 // indirect
	github.com/marten-seemann/qtls-go1-17 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.2 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.0-beta.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mholt/acmez v1.0.2 // indirect
	github.com/miekg/dns v1.1.50 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/oschwald/maxminddb-golang v1.9.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/pires/go-proxyproto v0.6.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.32.1 // indirect
	github.com/prometheus/procfs v0.7.3 // indirect
	github.com/refraction-networking/utls v1.1.0 // indirect
	github.com/riobard/go-bloom v0.0.0-20200614022211-cdc8013cb5b3 // indirect
	github.com/sagernet/sing v0.0.0-20220714145306-09b55ce4b6d0 // indirect
	github.com/sagernet/sing-shadowsocks v0.0.0-20220716012931-952ae62e05d7 // indirect
	github.com/seiflotfy/cuckoofilter v0.0.0-20220411075957-e3b120b3f5fb // indirect
	github.com/shadowsocks/go-shadowsocks2 v0.1.5 // indirect
	github.com/songgao/water v0.0.0-20190725173103-fd331bda3f4b // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.3.0 // indirect
	github.com/txthinking/runnergroup v0.0.0-20210608031112-152c7c4432bf // indirect
	github.com/txthinking/socks5 v0.0.0-20220212043548-414499347d4a // indirect
	github.com/txthinking/x v0.0.0-20210326105829-476fab902fbe // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/v2fly/ss-bloomring v0.0.0-20210312155135-28617310f63e // indirect
	github.com/v2fly/v2ray-core/v4 v4.42.1 // indirect
	github.com/xtaci/smux v1.5.15 // indirect
	github.com/xtls/go v0.0.0-20210920065950-d4af136d3672 // indirect
	go.starlark.net v0.0.0-20220714194419-4cadf0a12139 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
	golang.org/x/tools v0.1.11 // indirect
	google.golang.org/genproto v0.0.0-20220715211116-798f69b842b9 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/blake3 v1.1.7 // indirect
)

replace github.com/lucas-clemente/quic-go => github.com/tobyxdd/quic-go v0.28.1-0.20220706211558-7780039ad599
