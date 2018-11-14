package handler

import (
	"net/http"
	"strconv"

	"github.com/abasit/postwitter-api/model"

	"github.com/labstack/echo"
)

// GetUserPosts function
func GetUserPosts(c echo.Context) error {

	userID := c.Param("id")
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

	response := model.GetUserPosts(userID, initValue, limit)
	return c.JSON(http.StatusOK, response)
}

// GetUserPostFeed function
func GetUserPostFeed(c echo.Context) error {

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

	response := model.GetUserPostFeed(initValue, limit)
	return c.JSON(http.StatusOK, response)
}

// CreatePost function
func CreatePost(c echo.Context) (err error) {
	// Get User info
	u := &model.User{
		ID: userIDFromToken(c),
	}

	// Bind
	p := new(model.Post)
	if err = c.Bind(p); err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "unauthorized access"}
	}

	// Validate
	if p.Message == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid post message"}
	}

	// Create Post
	post := model.CreatePost(u.ID, p.Message)
	if post.ID <= 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Cannot post message at the moment"}
	}

	return c.JSON(http.StatusOK, p)
}
