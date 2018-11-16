package main

import (
	"os"
	"postwitter-api/route"

	"github.com/labstack/echo"
)

// @title Postwitter REST API
// @version 1.0
// @description This is a sample demo API server.
// @termsOfService https://postwitter-portal.herokuapp.com/terms/

// @contact.name API Support
// @contact.url https://postwitter-portal.herokuapp.com/support
// @contact.email abdulbasitmughal@gmail.com

// @license.name MIT 2.0
// @license.url https://opensource.org/licenses/MIT

// @host https://postwitter-portal.herokuapp.com
// @BasePath /v1

// @securityDefinitions.basic JWT

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

func main() {
	e := echo.New()

	// Routs
	router := route.Init()
	// Start server
	e.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))

}
