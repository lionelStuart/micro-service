package handler

import (
	"context"
	us "micro-service/user-srv/model/user"
	proto "micro-service/user-srv/proto/user"

	"github.com/micro/go-micro/util/log"
)

type Service struct{}

var (
	userService us.Service
)

func Init() {
	var err error

	userService, err = us.GetService()
	if err != nil {
		log.Fatal("[Init] init handler failed")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{Code: 500, Detail: err.Error()}
		return err
	}

	rsp.Success = true
	rsp.User = user
	return nil
}
