package payment

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"micro-service/basic/common"
	inv "micro-service/inventory-srv/proto/inventory"
	ord "micro-service/orders-srv/proto/order"
	"sync"
)

var (
	s            *service
	invClient    inv.InventoryService
	orderClient  ord.OrdersService
	m            sync.RWMutex
	payPublisher micro.Publisher
)

type service struct {
}

type Service interface {
	PayOrder(orderId int64) (err error)
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService uninited")
	}
	return s, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	invClient = inv.NewInventoryService("mu.micro.book.srv.inventory", client.DefaultClient)
	orderClient = ord.NewOrdersService("mu.micro.book.srv.orders", client.DefaultClient)
	payPublisher = micro.NewPublisher(common.TopicPaymentDone, client.DefaultClient)
	s = &service{}
}
