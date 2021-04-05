package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*Create mysql connection*/
func DB() *sql.DB {

	db, err := sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")"+DB_NAME+"?parseTime=true")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}

	//defer db.Close()

	return db
}
