package basic

import (
	"micro-service/user-srv/basic/config"
	"micro-service/user-srv/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
