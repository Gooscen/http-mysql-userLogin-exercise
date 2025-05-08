package main

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("数据库不可达:", err)
	}
}
