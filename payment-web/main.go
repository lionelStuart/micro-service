package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic"
	"micro-service/basic/config"
	"net/http"

	"github.com/micro/go-micro/web"
	"micro-service/payment-web/handler"
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

	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.payment"),
		web.Registry(micReg),
		web.Address(":8090"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()
		}),
	); err != nil {
		log.Fatal(err)
	}

	// register html handler
	authHandler := http.HandlerFunc(handler.PayOrder)
	service.Handle("/payment/pay-order", authHandler)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
