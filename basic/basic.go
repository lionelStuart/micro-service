package basic

import (
	"micro-service/basic/config"
	"micro-service/basic/db"
)

func Init() {
	config.Init()
	db.Init()
}
