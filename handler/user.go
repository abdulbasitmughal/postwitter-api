package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/abasit/postwitter-api/conf"

	"github.com/abasit/postwitter-api/model"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// GetUsers function
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

// GetUser function
func GetUser(c echo.Context) error {
	userID := c.Param("id")
	response := model.GetUser(userID)
	return c.JSON(http.StatusOK, response)
}

// Signup function
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
	user := model.CreateUser(u.Name, u.Email, u.Password)
	if user.ID <= 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid/unavailable email address"}
	}

	// Rest password
	u.Password = ""
	u.ID = user.ID

	return c.JSON(http.StatusOK, u)
}

// Login function
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
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(conf.KEY))
	if err != nil {
		return err
	}
	u.Name = user.Name
	u.TimeTag = user.TimeTag
	u.ID = user.ID
	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}

// userIDFromToken fetch userId from jwt token
func userIDFromToken(c echo.Context) int64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fmt.Println(claims["id"])

	id, err := strconv.Atoi(claims["id"].(string))

	fmt.Println(int64(id))

	if err != nil {
		fmt.Print(err)
	}
	return 123
}
