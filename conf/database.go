package conf

import (
	"database/sql"
	"github.com/prometheus/common/log"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:olid@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	return db
}