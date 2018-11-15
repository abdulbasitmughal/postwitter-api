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
	dbURIParts := strings.Split(dbURI, "//")[1]

	dbURISubParts := strings.Split(dbURIParts, "@")
	domain := dbURISubParts[1]
	dbURIPartAddress := strings.Split(domain, "/")
	dbName := dbURIPartAddress[1]
	dbNameString := strings.Split(dbName, "?")

	fmt.Println(dbURIParts[1])

	db, err := sql.Open("mysql", dbURISubParts[0]+"@tcp("+dbURIPartAddress[0]+":3306)/"+dbNameString[0])
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_twitter")
	if err != nil {
		panic(err)
	}

	//defer db.Close()
	// make sure connection is available
	// err = db.Ping()
	// fmt.Println(err)

	if err != nil {
		panic(err.Error)
	}

	return db
}
