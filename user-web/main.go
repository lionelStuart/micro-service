package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"
	"micro-service/basic/common"
	"micro-service/basic/config"
	"micro-service/plugins/hystrix"

	"micro-service/basic"

	"github.com/micro/go-micro/web"
	"micro-service/user-web/handler"
)

var (
	appName = "user_web"
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

	// create new web service
	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)

	// initialise service
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// register login handler
	service.HandleFunc("/user/login", handler.Login)
	service.HandleFunc("/user/logout", handler.Logout)

	// hystrix
	hystrix.StartStreamService("", "81")

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
