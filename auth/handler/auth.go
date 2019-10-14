package handler

import (
	"context"
	"micro-service/auth/model/access"

	"github.com/micro/go-micro/util/log"

	proto "micro-service/auth/proto/auth"
)

var (
	accessService access.Service
)

func Init() {
	var err error
	accessService, err = access.GetService()
	if err != nil {
		log.Fatal("[Init] init handle failure ,%s", err)
		return
	}
}

type Service struct {
}

func (s *Service) MakeAccessToken(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	panic("")
}

func (s *Service) DelUserAccessToken(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	panic("")
}
