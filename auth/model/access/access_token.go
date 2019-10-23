package access

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/util/log"
	"time"
)

var (
	//timeout = 30 day
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	tokenIDKeyPrefix  = "token:auth:id"
	tokenExpiredTopic = "mu.micro.book.topic.auth.tokenExpired"
)

type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

func (s *service) MakeAccessToken(subject *Subject) (ret string, err error) {
	// creat token claim
	m, err := s.createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] Create Token Claim Fail,err: %s", err)
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	ret, err = token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] Create Token Fail,err: %s", err)
	}

	// save token to cache: redis
	err = s.SaveTokenToCache(subject, ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] Save Token to Cache Fail,err: %s", err)
	}

	return
}

func (s *service) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] Load Token from Cache Fail,err:%s", err)
	}

	return

}

func (s *service) DelUserAccessToken(token string) (err error) {
	claims, err := s.parseToken(token)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] Wrong Token,err: %s", err)
	}

	// parse user id and del
	err = s.delTokenFromCache(&Subject{
		ID: claims.Subject,
	})
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] Del User Token Fail,err: %s", err)
	}

	// publish msg del
	msg := &broker.Message{
		Body: []byte(claims.Subject),
	}
	if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
		log.Logf("[pub] Publish Token Del Fail,err: %v", err)
	} else {
		fmt.Println("[pub] Publish Token Del Succ: ", string(msg.Body))
	}

	return
}
