package user

import (
	"github.com/go-log/log"
	"micro-service/plugins/db"
	proto "micro-service/user-srv/proto/user"
)

func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd FROM user WHERE user_name = ?`

	o := db.GetDB()
	ret = &proto.User{}

	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] fail query data,err %s", err)
		return
	}
	return
}
