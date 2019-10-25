package hystrix

import (
	hystrix_go "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/util/log"
	"net"
	"net/http"

	client2 "github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
)

func Init() {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1

}

func WrapperClient(c client2.Client) client2.Client {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	cl := hystrix.NewClientWrapper()(c)
	return cl
}

func StartStreamService(host, port string) (streamHandler *hystrix_go.StreamHandler) {
	streamHandler = hystrix_go.NewStreamHandler()
	streamHandler.Start()
	log.Logf("[Hystrix] Start Stream Service ...")
	go http.ListenAndServe(net.JoinHostPort(host, port), streamHandler)
	return
}
