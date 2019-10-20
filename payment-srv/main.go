package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic"
	"micro-service/basic/config"
	"micro-service/payment-srv/handler"
	"micro-service/payment-srv/model"

	payment "micro-service/payment-srv/proto/payment"
)

//complete registry options
func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Addrs = []string{
		fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort()),
	}
}

func main() {
	basic.Init()

	micReg := consul.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.payment"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(context *cli.Context) {
			model.Init()
			handler.Init()
		}),
	)

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
