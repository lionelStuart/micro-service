package access

import "time"

var (
	//timeout = 30 day
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	tokenIDKeyPrefix  = "token:auth:id"
	tokenExpiredTopic = "mu.micro.book.topic.auth.tokenExpired"
)

type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s service) MakeAccessToken(subject *Subject) (ret string, err error) {
	panic("implement me")
}

func (s service) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	panic("implement me")
}

func (s service) DelUserAccessToken(token string) (err error) {
	panic("implement me")
}
