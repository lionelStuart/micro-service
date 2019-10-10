package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"user-srv/handler"

	user "user-srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
