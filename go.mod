module github.com/cloustone/pandas

go 1.13

require (
	bou.ke/monkey v1.0.2
	github.com/Shopify/sarama v1.26.1
	github.com/bsm/sarama-cluster v2.1.15+incompatible
	github.com/carbocation/handlers v0.0.0-20140528190747-c939c6d9ef31 // indirect
	github.com/carbocation/interpose v0.0.0-20161206215253-723534742ba3
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/didip/tollbooth v4.0.2+incompatible
	github.com/dre1080/recover v0.0.0-20150930082637-1c296bbb3227
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/fatih/structs v1.1.0
	github.com/garyburd/redigo v1.6.0
	github.com/go-ldap/ldap v3.0.3+incompatible
	github.com/go-macaron/binding v1.1.0
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab // indirect
	github.com/go-openapi/errors v0.19.4
	github.com/go-openapi/loads v0.19.5
	github.com/go-openapi/runtime v0.19.12
	github.com/go-openapi/spec v0.19.7
	github.com/go-openapi/strfmt v0.19.5
	github.com/go-openapi/swag v0.19.8
	github.com/go-openapi/validate v0.19.7
	github.com/gogo/protobuf v1.3.1
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.5
	github.com/goods/httpbuf v0.0.0-20120503183857-5709e9bb814c // indirect
	github.com/interpose/middleware v0.0.0-20150216143757-05ed56ed52fa // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/jinzhu/gorm v1.9.12
	github.com/justinas/nosurf v1.1.0 // indirect
	github.com/meatballhat/negroni-logrus v1.1.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/phyber/negroni-gzip v0.0.0-20180113114010-ef6356a5d029 // indirect
	github.com/rs/xid v1.2.1
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	google.golang.org/grpc v1.26.0
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/macaron.v1 v1.3.5
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/apimachinery v0.17.4
)

replace gopkg.in/asn1-ber.v1 => github.com/go-asn1-ber/asn1-ber v1.4.1
