package db

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	// load drivers for mysql
	_ "github.com/go-sql-driver/mysql"
)

// CreateCon mysql database connection
func CreateCon() *sql.DB {

	dbURI := os.Getenv("CLEARDB_DATABASE_URL")
	dbURIParts := strings.Split(dbURI, "//")

	db, err := sql.Open("mysql", dbURIParts[1])
	if err != nil {
		panic(err)
	}

	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)

	if err != nil {
		panic(err)
	}

	return db
}
