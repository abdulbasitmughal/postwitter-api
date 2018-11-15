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

// ResponsePost struct
type ResponsePost struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Message string `json: "message"`
	TimeTag string `json : "time_tag"`
}

// Posts struct
type Posts struct {
	Posts []Post `json:"post"`
}

// ResponsePosts struct
type ResponsePosts struct {
	ResponsePosts []ResponsePost `json:"post"`
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
func GetUserPostFeed(initValue int, limit int) ResponsePosts {
	con := db.CreateCon()
	//db.CreateCon()
	sqlStatement := fmt.Sprintf("SELECT user.email, post.message, post.time_tag FROM post INNER JOIN user ON user.id = post.user_id ORDER BY post.time_tag desc LIMIT %d,%d", initValue, limit)

	rows, err := con.Query(sqlStatement)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	result := ResponsePosts{}

	for rows.Next() {
		post := ResponsePost{}

		err2 := rows.Scan(&post.Email, &post.Message, &post.TimeTag)
		// Exit if we get an error
		if err2 != nil {
			fmt.Print(err2)
		}

		result.ResponsePosts = append(result.ResponsePosts, post)
	}

	return result
}

// GetPostByID function
func GetPostByID(id int64) ResponsePost {
	con := db.CreateCon()

	p := ResponsePost{}
	err := con.QueryRow("SELECT post.id, message, post.time_tag, email FROM post INNER JOIN user ON user.id = post.user_id WHERE post.id = ?", id).Scan(&p.ID, &p.Message, &p.TimeTag, &p.Email)

	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}
	return p
}

// CreatePost function
func CreatePost(email string, message string) ResponsePost {
	con := db.CreateCon()

	p := ResponsePost{}
	user := GetUserByEmail(email)

	res, err := con.Exec("INSERT INTO post (user_id, message) VALUES (?, ?)", user.ID, message)

	if err != nil {
		p.ID = -1
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
		} else {
			p.ID = id
		}
	}

	defer con.Close()

	return GetPostByID(p.ID)
}
