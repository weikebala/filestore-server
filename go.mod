module filestore-server

go 1.14

require (
	github.com/aliyun/aliyun-oss-go-sdk v2.1.4+incompatible
	github.com/garyburd/redigo v1.6.2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/contrib v0.0.0-20200810162008-6dee08bf958e
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/gomodule/redigo v1.8.2
	github.com/json-iterator/go v1.1.9
	github.com/juju/ratelimit v1.0.1
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/breaker/hystrix v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/wrapper/ratelimiter/ratelimit v0.0.0-20200119172437-4fe21aa238fd
	github.com/mitchellh/mapstructure v1.1.2
	github.com/moxiaomomo/go-bindata-assetfs v1.0.0
	github.com/streadway/amqp v1.0.0
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gopkg.in/amz.v1 v1.0.0-20150111123259-ad23e96a31d2
)
