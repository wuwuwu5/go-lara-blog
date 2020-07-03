package models

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

type DB struct {
}

var db *sql.DB

func InitDB() {
	var err error

	db, err = sql.Open("mysql", "admin:q1w2e3r4@tcp(192.168.10.10)/blog?charset=utf8mb4")

	if err != nil {
		log.Fatal("数据库参数失败")
		return
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("数据库连接失败")
		return
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(50)
}

func GetDB() *sql.DB {
	return db
}
