module github.com/PCBismarck/tiktok_server/cmd/api

go 1.18

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/PCBismarck/tiktok_server/cmd/comment v0.0.0-20230223035059-2525bbc54cfe
	github.com/PCBismarck/tiktok_server/cmd/favorite v0.0.0-20230223134202-cca779328b6e
	github.com/PCBismarck/tiktok_server/cmd/relation v0.0.0-20230223134202-cca779328b6e
	github.com/PCBismarck/tiktok_server/cmd/user v0.0.0-20230223134202-cca779328b6e
	github.com/apache/thrift v0.18.0
	github.com/cloudwego/hertz v0.5.2
	github.com/cloudwego/kitex v0.4.4
)

require github.com/golang-jwt/jwt/v4 v4.4.1 // indirect

require (
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.0.0-20221122125632-68358b8ecec6 // indirect
	github.com/bytedance/sonic v1.5.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20211019084208-fb5309c8db06 // indirect
	github.com/chenzhuoyu/iasm v0.0.0-20230222070914-0b1b64b0e762 // indirect
	github.com/choleraehyq/pid v0.0.16 // indirect
	github.com/cloudwego/fastpb v0.0.3 // indirect
	github.com/cloudwego/frugal v0.1.5 // indirect
	github.com/cloudwego/netpoll v0.3.2 // indirect
	github.com/cloudwego/thriftgo v0.2.7 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/pprof v0.0.0-20230222194610-99052d3372e7 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/hertz-contrib/jwt v1.0.2
	github.com/jhump/protoreflect v1.14.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/oleiade/lane v1.0.1 // indirect
	github.com/tidwall/gjson v1.14.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.2.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230222225845-10f96fb3dbec // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
