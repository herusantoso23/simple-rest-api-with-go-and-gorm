package services

import (
	"net/http"
	"simple-rest-api-using-gorm-and-echo/handler"
	"simple-rest-api-using-gorm-and-echo/structs"

	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) (err error) {
	u := new(structs.User)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		return
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse(u))
}

func UpdateUser(c echo.Context) (err error) {
	u := new(structs.User)
	if err = c.Bind(u); err != nil {
		return
	}

	if u.ID == "" {
		return c.JSON(http.StatusBadRequest, handler.ErrorResponse(http.StatusBadRequest, "ID must not be null", nil))
	} else {
		if err = c.Validate(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, handler.SuccessResponse(u))
	}
}

func GetUser(c echo.Context) (err error) {
	id := c.Param("id")

	u := structs.User{
		ID:      id,
		Name:    "Heru Santoso",
		Address: "Jakarta",
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse(u))
}

func GetAllUser(c echo.Context) (err error) {
	users := []structs.User{
		{
			ID:      "1",
			Name:    "Heru",
			Address: "Jakarta",
		},
		{
			ID:      "2",
			Name:    "Heru Santoso",
			Address: "Jakarta",
		},
		{
			ID:      "3",
			Name:    "Santoso Heru",
			Address: "Jakarta",
		},
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse(users))
}
