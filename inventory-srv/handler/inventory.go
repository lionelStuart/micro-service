package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	inv "micro-service/inventory-srv/model/inventory"
	proto "micro-service/inventory-srv/proto/inventory"
)

var (
	invService inv.Service
)

type Service struct {
}

func Init() {
	invService, _ = inv.GetService()
}

func (e *Service) Sell(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	id, err := invService.Sell(req.BookId, req.UserId)
	if err != nil {
		log.Logf("[SELL] sell failed, bookId: %d, userId: %d, %s", req.BookId, req.UserId, err)
		return
	}

	rsp.InvH = &proto.InvHistory{Id: id}
	rsp.Success = true

	return nil
}

func (e *Service) Confirm(ctx context.Context, req *proto.Request, rsp *proto.Response) (err error) {
	err = invService.Confirm(req.HistoryId, int(req.HistoryState))
	if err != nil {
		log.Logf("[Confirm] confirm sell failed, %s", err)
		return
	}

	rsp.Success = true
	return nil
}
