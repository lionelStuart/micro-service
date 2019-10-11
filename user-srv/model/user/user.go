package user

import (
	"fmt"
	"sync"
	proto "user-srv/proto/user"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

type Service interface {
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// get service instance
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] not inited ")
	}

	return s, nil

}

//init user service
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	s = &service{}

}
