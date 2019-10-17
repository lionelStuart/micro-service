package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic"
	"micro-service/basic/common"
	"micro-service/basic/config"
	"micro-service/orders-srv/handler"
	"micro-service/orders-srv/subscriber"
	"micro-service/user-srv/model"

	order "micro-service/orders-srv/proto/order"
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

	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.order"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()
			handler.Init()
			subscriber.Init()
		}),
	)

	// Register Handler
	order.RegisterOrdersHandler(service.Server(), new(handler.Order))

	// Register PaySub
	err := micro.RegisterSubscriber(common.TopicPaymentDone, service.Server(), subscriber.PayOrder)
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
