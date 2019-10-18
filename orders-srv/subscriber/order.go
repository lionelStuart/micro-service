package subscriber

import (
	"context"
	_ "github.com/micro/go-micro/util/log"

	order "micro-service/orders-srv/model/order"
)

var (
	orderService order.Service
)

func Init() {
	orderService, _ = order.GetService()
}

func PayOrder(ctx context.Context, un interface{}) (err error) {
	panic("TODO")
}
