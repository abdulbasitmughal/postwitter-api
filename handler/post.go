package handler

import (
	"net/http"
	"strconv"

	"postwitter-api/model"

	"github.com/labstack/echo"
)

// GetUserPosts godoc
// @Summary List Users Posts
// @Description get list of user posts using email address
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Post
// @Router /v1/users/:email/posts [get]
func GetUserPosts(c echo.Context) error {

	email := c.Param("email")
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

	response := model.GetUserPosts(email, initValue, limit)
	return c.JSON(http.StatusOK, response)
}

// GetUserPostFeed godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "page search by q"
// @Success 200 {array} model.Post
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Router /v1/posts [get]
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

// CreatePost godoc
// @Summary Create Post
// @Description get accounts
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Post
// @Failure 400 {object} echo.HTTPError
// @Failure 401 {object} echo.HTTPError
// @Router /v1/posts [post]
func CreatePost(c echo.Context) (err error) {
	// Get User info
	u := &model.User{
		Email: userIDFromToken(c),
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
	post := model.CreatePost(u.Email, p.Message)
	if post.ID <= 0 {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "Cannot post message at the moment"}
	}

	return c.JSON(http.StatusOK, post)
}
