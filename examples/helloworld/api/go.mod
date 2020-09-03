module api

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/xxdawn/micro/examples/helloworld/api => github.com/micro/services/helloworld/api v0.0.0-20200706113312-8a497d4e1aaa

require (
	github.com/xxdawn/micro/examples/helloworld v0.0.0-20200706113312-8a497d4e1aaa
	github.com/micro/go-micro/v2 v2.9.1
)
