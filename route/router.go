package route

import (
	"postwitter-api/conf"
	"postwitter-api/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/swaggo/echo-swagger"
)

// Init the routs
func Init() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// TODO Move middleware in middleware folder
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// TODO Move middleware in middleware folder
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf.KEY),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/v1/login" || c.Path() == "/v1/signup" || c.Path() == "/v1/swagger" {
				return true
			}
			return false
		},
	}))

	// Routes
	// V1 API's
	v1 := e.Group("/v1")
	{
		// Routes Auth
		v1.POST("/signup", handler.Signup)
		v1.POST("/login", handler.Login)

		// User Routs
		v1.GET("/users", handler.GetUsers)
		v1.GET("/users/:email/posts", handler.GetUserPosts)

		// Routes Posts
		v1.GET("/posts", handler.GetUserPostFeed)
		v1.POST("/posts", handler.CreatePost)

		// Routes Swagger
		v1.GET("/swagger/*", echoSwagger.WrapHandler)
	}

	return e
}
