package main

import (
	"os"

	"github.com/abasit/postwitter-api/handler"

	"github.com/abasit/postwitter-api/conf"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.KEY),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	}))

	// Routes
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	e.GET("/users", handler.GetUsers)
	e.GET("/users/:id", handler.GetUser)
	e.GET("/users/:id/posts", handler.GetUserPosts)
	e.GET("/posts", handler.GetUserPostFeed)

	e.POST("/posts", handler.CreatePost)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
