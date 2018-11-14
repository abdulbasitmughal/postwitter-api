package db

import (
	"database/sql"
	"fmt"

	"postwitter-api/conf"
	// load drivers for mysql
	_ "github.com/go-sql-driver/mysql"
)

// CreateCon mysql database connection
func CreateCon() *sql.DB {

	db, err := sql.Open("mysql", conf.DBUSER+":"+conf.DBPASSWORD+"@tcp("+conf.DBHOST+":"+conf.DBPORT+")/"+conf.DBNAME)
	if err != nil {
		fmt.Println(err.Error())
	}

	//defer db.Close()
	// make sure connection is available
	// err = db.Ping()
	// fmt.Println(err)

	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}

	return db
}
