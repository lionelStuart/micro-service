package basic

import (
	"micro-service/basic/config"
	"micro-service/basic/db"
	"micro-service/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
