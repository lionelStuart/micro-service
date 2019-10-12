package db

import (
	"database/sql"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic/config"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func Init() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db inited already")
		log.Logf(err.Error())
		return
	}

	// init mysql if enabled
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
