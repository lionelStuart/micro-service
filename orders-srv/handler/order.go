package handler

import (
	"context"
	"github.com/go-log/log"
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
	orderId, err := orderService.New(req.BookId, req.UserId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{Detail: err.Error()}
		return
	}

	rsp.Order = &proto.Order{Id: orderId}
	return
}

func (e *Order) GetOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Logf("[GetOrder] Recv New Order Request, %d", req.OrderId)

	rsp.Order, err = orderService.GetOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{Detail: err.Error()}
		return
	}

	rsp.Success = true
	return
}
