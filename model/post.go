package model

import (
	"fmt"

	"postwitter-api/db"
)

// Post struct mapped over database table
type Post struct {
	ID      int64  `json:"id"`
	UserID  User   `json:"user_id"`
	Message string `json: "message"`
	TimeTag string `json : "time_tag"`
}

// Posts struct
type Posts struct {
	Posts []Post `json:"post"`
}

// GetUserPosts function
func GetUserPosts(email string, initValue int, limit int) Posts {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := fmt.Sprintf("SELECT id, message, time_tag FROM post INNER JOIN user ON user.id = post.user_id WHERE email = %s ORDER BY time_tag desc LIMIT %d,%d", email, initValue, limit)

	rows, err := con.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Posts{}

	for rows.Next() {
		post := Post{}

		err2 := rows.Scan(&post.ID, &post.Message, &post.TimeTag)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}

		result.Posts = append(result.Posts, post)
	}

	return result
}

// GetUserPostFeed function
func GetUserPostFeed(initValue int, limit int) Posts {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := fmt.Sprintf("SELECT id, message, time_tag FROM post ORDER BY time_tag desc LIMIT %d,%d", initValue, limit)

	rows, err := con.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := Posts{}

	for rows.Next() {
		post := Post{}

		err2 := rows.Scan(&post.ID, &post.Message, &post.TimeTag)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}

		result.Posts = append(result.Posts, post)
	}

	return result
}

// CreatePost function
func CreatePost(email string, message string) Post {
	con := db.CreateCon()

	p := Post{}
	userID := GetUserByEmail(email)
	res, err := con.Exec("INSERT INTO post (user_id, message) VALUES (?, ?)", userID, message)

	if err != nil {
		p.ID = -1
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
		} else {
			p.ID = id
			println("LastInsertId:", id)
		}
	}

	defer con.Close()

	return p
}
