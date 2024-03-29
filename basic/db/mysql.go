package db

import (
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic/config"
)

func initMysql() {
	var err error

	// open mysql on conf url
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	//max idle connections
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	//max open connections
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
