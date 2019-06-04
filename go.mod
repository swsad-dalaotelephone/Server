module baobaozhuan

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go latest
	golang.org/x/build => github.com/golang/build latest
	golang.org/x/crypto => github.com/golang/crypto latest
	golang.org/x/exp => github.com/golang/exp latest
	golang.org/x/lint => github.com/golang/lint latest
	golang.org/x/net => github.com/golang/net latest
	golang.org/x/oauth2 => github.com/golang/oauth2 latest
	golang.org/x/perf => github.com/golang/perf latest
	golang.org/x/sync => github.com/golang/sync latest
	golang.org/x/sys => github.com/golang/sys latest
	golang.org/x/text => github.com/golang/text latestl
	golang.org/x/time => github.com/golang/time latest
	golang.org/x/tools => github.com/golang/tools latest
	google.golang.org/api => github.com/google/google-api-go-client latest
	google.golang.org/appengine => github.com/golang/appengine latest
	google.golang.org/genproto => github.com/google/go-genproto latest
	google.golang.org/grpc => github.com/grpc/grpc-go latest
)

require (
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/gin-contrib/cache v1.1.0
	github.com/gin-contrib/logger v0.0.1
	github.com/gin-contrib/sessions v0.0.0-20190226023029-1532893d996f
	github.com/gin-gonic/gin v1.3.0
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/jinzhu/gorm v1.9.4
	github.com/jinzhu/now v1.0.0 // indirect
	github.com/medivhzhan/weapp v1.1.1
	github.com/rs/zerolog v1.14.3
	github.com/satori/go.uuid v1.2.0
)
