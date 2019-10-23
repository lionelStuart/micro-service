package db

import (
	"database/sql"
	"github.com/micro/go-micro/util/log"
	"micro-service/basic/config"
)

type db struct {
	Mysql mysql `json:"mysql"`
}

type mysql struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enabled"`
	maxIdleConnection int    `json:"maxIdleConnection"`
	maxOpenConnection int    `json:"maxOpenConnection"`
}

func initMysql() {
	var err error

	c := config.C()
	cfg := &db{}

	err = c.App("db", cfg)
	if err != nil {
		log.Logf("[initMysql] %s", err)
	}

	if !cfg.Mysql.Enable {
		log.Log("[initMysql] mysql not inited")
	}

	// open mysql on conf url
	mysqlDB, err = sql.Open("mysql", cfg.Mysql.URL)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	//max idle connections
	mysqlDB.SetMaxIdleConns(cfg.Mysql.maxIdleConnection)

	//max open connections
	mysqlDB.SetMaxOpenConns(cfg.Mysql.maxOpenConnection)

	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Logf("[initMysql] connection mysql success")
}
