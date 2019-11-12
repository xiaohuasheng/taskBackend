package db_utils

import (
	. "config"
	. "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var SqlDB *DB

type conf struct {
	Enabled bool   `yaml:"enabled"` //yaml：yaml格式 enabled：属性的为enabled
	Path    string `yaml:"path"`
}

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
	var c Config
	c.GetConf()

	env := os.Getenv("ENV")
	var dataSourceName string
	if "dev" == env {
		dataSourceName = "root:123456@tcp(192.168.205.3:3306)/task?tls=skip-verify&autocommit=true"
	} else {
		dataSourceName = fmt.Sprintf("%s:%s@(%s:%d)/task", c.DBConfig.User, c.DBConfig.Password, c.DBConfig.Host, c.DBConfig.Port)
	}
	fmt.Println(dataSourceName)
	SqlDB, err := Open("mysql", dataSourceName)
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
