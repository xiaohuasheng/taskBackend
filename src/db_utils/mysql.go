package db_utils

import (
	. "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var SqlDB *DB

func GetDB() *DB {
	if SqlDB != nil {
		return SqlDB
	}
	SqlDB = initDB()
	fmt.Println(SqlDB)
	return SqlDB
}

func initDB() *DB {
	fmt.Println("开始初始化...")
	SqlDB, err := Open("mysql", "root:123456@tcp(192.168.205.3:3306)/task?tls=skip-verify&autocommit=true")
	if err != nil {
		log.Fatalln(err)
	}

	SqlDB.SetMaxIdleConns(20)
	SqlDB.SetMaxOpenConns(20)

	if err := SqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("初始化结束...")
	return SqlDB
}
