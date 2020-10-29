package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/todo_dev?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Printf("%#v\n", err)
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
