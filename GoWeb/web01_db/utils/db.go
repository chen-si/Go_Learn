package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)
//SQL 结构化查询语言
func init() {
	Db, err = sql.Open("mysql","root:1234w5asd@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
}
