package subscriber

import (
	"context"
	"github.com/go-log/log"
	_ "github.com/micro/go-micro/util/log"

	order "micro-service/orders-srv/model/order"
	payment "micro-service/payment-srv/proto/payment"
)

var (
	orderService order.Service
)

func Init() {
	orderService, _ = order.GetService()
}

func PayOrder(ctx context.Context, event *payment.PayEvent) (err error) {
	log.Logf("[PayOrder] Recv Order Payment Sub, %d, %d", event.OrderId, event.State)

	err = orderService.UpdateOrderState(event.OrderId, int(event.State))
	if err != nil {
		log.Logf("[PayOrder] Recv Pay sub, update state err, %s", err)
		return
	}

	return
}
