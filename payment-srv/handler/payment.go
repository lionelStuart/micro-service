package handler

import (
	"context"
	"github.com/micro/go-micro/util/log"

	"micro-service/payment-srv/model/payment"
	proto "micro-service/payment-srv/proto/payment"
)

var (
	paymentService payment.Service
)

type Service struct {
}

func Init() {
	paymentService, _ = payment.GetService()
}

func (e *Service) PayOrder(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	log.Log("[PayOrder] Recv Request")

	err = paymentService.PayOrder(req.OrderId)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Detail: err.Error()}
	}
	rsp.Success = true
	return
}
