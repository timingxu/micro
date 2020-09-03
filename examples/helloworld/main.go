package main

import (
	"github.com/xxdawn/micro/examples/helloworld/handler"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"

	helloworld "github.com/xxdawn/micro/examples/helloworld/proto"
	consul "github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// New Service

	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.helloworld"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	helloworld.RegisterHelloworldHandler(service.Server(), new(handler.Helloworld))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
