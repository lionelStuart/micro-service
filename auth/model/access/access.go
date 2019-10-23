package access

import (
	"fmt"
	r "github.com/go-redis/redis"
	"micro-service/basic/config"
	"micro-service/plugins/jwt"
	"micro-service/plugins/redis"
	"sync"
)

var (
	s   *service
	ca  *r.Client
	m   sync.RWMutex
	cfg = &jwt.Jwt{}
)

type service struct {
}

type Service interface {
	MakeAccessToken(subject *Subject) (ret string, err error)
	GetCachedAccessToken(subject *Subject) (ret string, err error)
	DelUserAccessToken(token string) (err error)
}

func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	err := config.C().App("jwt", cfg)
	if err != nil {
		panic(err)
	}

	ca = redis.GetRedis()

	s = &service{}
}

func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] service not inited")
	}
	return s, nil
}
