package database

import (
	"database/sql"
	"log"
)
var Db *sql.DB
func ClientDB() {
        dsn := "root:114514@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
        
	db, err := sql.Open("message_board", dsn)
        if err != nil {
                log.Fatal(err)
        }
        err = db.Ping()
        if err != nil {
		log.Fatal(err)
                       }
	Db=db
}

