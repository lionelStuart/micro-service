package payment

import (
	"context"
	"github.com/go-log/log"
	"github.com/google/uuid"
	proto "micro-service/payment-srv/proto/payment"
	"time"
)

func (s *service) sendPayDoneEvt(orderId int64, state int32) {
	evt := &proto.PayEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		OrderId:  orderId,
		State:    state,
	}

	log.Logf("[sendPayEvt] payEvt, %v \n", evt)

	if err := payPublisher.Publish(context.Background(), evt); err != nil {
		log.Logf("[SendPayDoneEvt] err :%v \n", err)
	}
}
