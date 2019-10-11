package user

import (
	"github.com/go-log/log"
	"user-srv/basic/db"
	proto "user-srv/proto/user"
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
