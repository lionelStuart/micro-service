package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"inventory-srv/handler"
	"inventory-srv/subscriber"

	inventory "inventory-srv/proto/inventory"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.inventory"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	inventory.RegisterInventoryHandler(service.Server(), new(handler.Inventory))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.srv.inventory", service.Server(), new(subscriber.Inventory))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.book.srv.inventory", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
