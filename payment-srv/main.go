package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"payment-srv/handler"
	"payment-srv/subscriber"

	payment "payment-srv/proto/payment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.payment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), new(handler.Payment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.srv.payment", service.Server(), new(subscriber.Payment))

	// Register Function as Subscriber
	micro.RegisterSubscriber("mu.micro.book.srv.payment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
