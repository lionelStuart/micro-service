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
	_ "net/http"

	"github.com/micro/go-micro/web"
	"micro-service/orders-web/handler"
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

	micReg := consul.NewRegistry()

	// create new web service
	service := web.NewService(
		web.Name("mu.micro.book.web.orders"),
		web.Registry(micReg),
		web.Address(":8091"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(
		web.Action(
			func(context *cli.Context) {
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	authHandler := http.HandlerFunc(handler.New)
	service.Handle("/orders/new", handler.AuthWrapper(authHandler))
	service.HandleFunc("/hi", handler.Hello)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
