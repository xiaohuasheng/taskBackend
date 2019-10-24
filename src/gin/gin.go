package gin

import (
	. "db_utils"
	_ "github.com/go-sql-driver/mysql"
	. "router"
)

func WebMain() {
	db := GetDB()
	defer db.Close()
	r := InitRouter()
	r.Run(":9999")
}
