package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"
	"micro-service/basic"
	"micro-service/basic/common"
	"micro-service/basic/config"
	"micro-service/orders-srv/handler"
	"micro-service/orders-srv/model"
	"micro-service/orders-srv/subscriber"

	order "micro-service/orders-srv/proto/order"
)

var (
	appName = "orders_srv"
	cfg     = &userCfg{}
)

type userCfg struct {
	common.AppCfg
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)
	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] ,%v", cfg)
	return
}

//complete registry options
func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{
		fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port),
	}

}

func main() {
	//init conf, sql ...
	initCfg()

	micReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
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
