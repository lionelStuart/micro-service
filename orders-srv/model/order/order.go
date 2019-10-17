package order

import (
	"fmt"
	"github.com/micro/go-micro/client"
	inv "micro-service/inventory-srv/proto/inventory"
	proto "micro-service/orders-srv/proto/order"
	"sync"
)

var (
	s         *service
	invClient inv.InventoryService
	m         sync.RWMutex
)

type service struct {
}

type Service interface {
	New(bookId, userId int64) (orderId int64, err error)
	GetOrder(orderId int64) (order *proto.Order, err error)
	UpdateOrderState(orderId int64, state int) (err error)
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
	s = &service{}
}
