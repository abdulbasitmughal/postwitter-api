package handler

import (
	"net/http"
	"strconv"
	"time"

	"postwitter-api/conf"

	"postwitter-api/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetUsers function
// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts/{id} [get]
func GetUsers(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	// Defaults
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 100
	}

	initValue := (page - 1) * limit

	response := model.GetUsers(initValue, limit)
	return c.JSON(http.StatusOK, response)
}

// Signup function
// Login godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [get]
func Signup(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Create User
	user := model.CreateUser(u.Email, u.Password)
	if user.ID <= 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid/unavailable email address"}
	}

	// Rest password
	u.Password = ""
	u.ID = user.ID

	return c.JSON(http.StatusOK, u)
}

// Login godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /accounts [get]
func Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	user := model.ValidateUser(u.Email, u.Password)
	if user.ID <= 0 {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(conf.KEY))
	if err != nil {
		return err
	}
	u.TimeTag = user.TimeTag
	u.ID = user.ID
	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}

// userIDFromToken fetch userId from jwt token
func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["email"].(string)
}
