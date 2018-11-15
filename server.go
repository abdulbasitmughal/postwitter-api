package main

import (
	"os"
	"postwitter-api/handler"

	"postwitter-api/conf"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/swaggo/echo-swagger"
)

// @title Postwitter REST API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

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
			if c.Path() == "/v1/login" || c.Path() == "/v1/signup" || c.Path() == "/v1/swagger" {
				return true
			}
			return false
		},
	}))

	e.GET("/v1/swagger/*", echoSwagger.WrapHandler)

	// Routes
	e.POST("/v1/signup", handler.Signup)
	e.POST("/v1/login", handler.Login)

	e.GET("/v1/users", handler.GetUsers)
	e.GET("/v1/users/:email/posts", handler.GetUserPosts)

	e.GET("/v1/posts", handler.GetUserPostFeed)
	e.POST("/v1/posts", handler.CreatePost)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
