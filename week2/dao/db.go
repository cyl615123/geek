package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

var db *sql.DB

func InitDB() error {
	dns := ""
	mysqlDB, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	db = mysqlDB
	return nil
}
