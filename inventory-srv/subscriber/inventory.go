package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	inventory "inventory-srv/proto/inventory"
)

type Inventory struct{}

func (e *Inventory) Handle(ctx context.Context, msg *inventory.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *inventory.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
