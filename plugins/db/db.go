package db

import (
	"database/sql"
	"fmt"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic"
	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

func init() {
	basic.Register(initDB)
}

func initDB() {
	m.Lock()
	defer m.Unlock()

	var err error

	if inited {
		err = fmt.Errorf("[Init] db inited already")
		log.Logf(err.Error())
		return
	}

	// init mysql if enabled
	initMysql()

	inited = true
}

func GetDB() *sql.DB {
	return mysqlDB
}
