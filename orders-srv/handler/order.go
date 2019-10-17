package handler

import (
	"context"
	"micro-service/orders-srv/model/order"

	proto "micro-service/orders-srv/proto/order"
)

var (
	orderService order.Service
)

type Order struct{}

func Init() {
	orderService, _ = order.GetService()
}

func (e *Order) New(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	panic("")
}

func (e *Order) GetOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	panic("")
}
