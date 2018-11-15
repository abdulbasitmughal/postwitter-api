package model

import (
	"database/sql"
	"fmt"

	"postwitter-api/db"

	// _ "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// User struct mapped over database table
type User struct {
	ID       int64  `db:"id" json:"id"`
	Email    string `db:"email" json: "email"`
	Password string `db:"password" json: "password"`
	Token    string `db:"token" json: "token"`
	TimeTag  string `json : "time_tag"`
}

// Users struct
type Users struct {
	Users []User `json:"user"`
}

var con *sql.DB

// GetUsers function
func GetUsers(initValue int, limit int) Users {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := fmt.Sprintf("SELECT id, email, time_tag FROM user ORDER BY email LIMIT %d,%d", initValue, limit)

	rows, err := con.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Users{}

	for rows.Next() {
		user := User{}

		err2 := rows.Scan(&user.ID, &user.Email, &user.TimeTag)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}

		result.Users = append(result.Users, user)
	}

	return result
}

// ValidateUser function
func ValidateUser(email string, password string) User {
	con := db.CreateCon()
	//db.CreateCon()
	u := User{}
	err := con.QueryRow("SELECT id, time_tag FROM user WHERE email = ? AND password = ?", email, password).Scan(&u.ID, &u.TimeTag)

	defer con.Close()

	if err != nil {
		fmt.Println(err)
	}

	return u
}

// CreateUser function
func CreateUser(email string, password string) User {
	con := db.CreateCon()

	u := User{}
	err := con.QueryRow("SELECT id, time_tag FROM user WHERE email = ?", email).Scan(&u.ID, &u.TimeTag)

	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("user exist:", u.ID)

	if u.ID > 0 {
		fmt.Println("user exist:", u.ID)
		u.ID = 0
	} else {
		res, err := con.Exec("INSERT INTO user (email, password) VALUES (?, ?)", email, password)

		if err != nil {
			u.ID = -1
		} else {
			id, err := res.LastInsertId()
			if err != nil {
				println("Error:", err.Error())
			} else {
				u.ID = id
				println("LastInsertId:", id)
			}
		}
	}

	defer con.Close()

	return u
}

// GetUserByEmail function
func GetUserByEmail(email string) User {
	con := db.CreateCon()

	u := User{}
	err := con.QueryRow("SELECT id, time_tag FROM user WHERE email = ?", email).Scan(&u.ID, &u.TimeTag)

	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}

	return u
}
