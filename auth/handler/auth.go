package handler

import (
	"context"
	"micro-service/auth/model/access"
	"strconv"

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
	log.Log("[MakeAccessToken] Recv Token Create Request")

	// make accessToken
	token, err := accessService.MakeAccessToken(&access.Subject{
		ID:   strconv.FormatUint(req.UserId, 10),
		Name: req.UserName,
	})
	if err != nil {
		rsp.Error = &proto.Error{Detail: err.Error()}
		log.Logf("[MakeAccessToken] Token Generate Fail,err: %s", err)
		return err
	}

	rsp.Token = token
	return nil
}

func (s *Service) DelUserAccessToken(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Log("[DelUserAccessToken] Delete User Token")
	err := accessService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &proto.Error{Detail: err.Error()}
		log.Logf("[MakeAccessToken] Token Del Fail,err: %s", err)
		return err
	}

	return nil

}
