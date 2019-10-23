package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-plugins/config/source/grpc"
	"micro-service/basic/common"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"micro-service/auth/handler"
	"micro-service/auth/model"
	"micro-service/basic"
	"micro-service/basic/config"

	auth "micro-service/auth/proto/auth"
)

var (
	appName = "auth_srv"
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

	micReg := consul.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
	)

	// Initialise service
	service.Init(
		micro.Action(func(ctx *cli.Context) {
			model.Init()
			handler.Init()
		}))

	// Register Handler
	auth.RegisterServiceHandler(service.Server(), &handler.Service{})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
