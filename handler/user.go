package handler

import (
	"net/http"
	"strconv"

	"postwitter-api/model"

	"github.com/labstack/echo"
)

// GetUsers godoc
// @Summary Get users list registered into the system.
// @Description get user list with page number and limit for page
// @ID get-string
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Account
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /v1/users [get]
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
