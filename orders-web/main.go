package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"
	"micro-service/basic"
	"micro-service/basic/common"
	"micro-service/basic/config"
	"net/http"
	_ "net/http"

	"github.com/micro/go-micro/web"
	"micro-service/orders-web/handler"
)

var (
	appName = "orders_web"
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

	// create new web service
	service := web.NewService(
		web.Name(cfg.Name),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
		web.Version(cfg.Version),
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
