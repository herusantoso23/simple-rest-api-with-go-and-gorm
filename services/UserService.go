package services

import (
	"net/http"
	"simple-rest-api-using-gorm-and-echo/handler"
	"simple-rest-api-using-gorm-and-echo/structs"

	"github.com/labstack/echo"
)

func (idb *InDB) CreateUser(c echo.Context) (err error) {
	u := new(structs.User)
	if err = c.Bind(u); err != nil {
		return
	}
	if err = c.Validate(u); err != nil {
		return
	}

	idb.DB.Create(&u)
	return c.JSON(http.StatusOK, handler.SuccessResponse(u))
}

func (idb *InDB) UpdateUser(c echo.Context) (err error) {
	dto := new(structs.User)
	if err = c.Bind(dto); err != nil {
		return
	}

	// Check user on database
	var u structs.User
	errs := idb.DB.First(&u, dto.ID).Error
	if errs != nil {
		return c.JSON(http.StatusNotFound, handler.ErrorResponse(http.StatusNotFound, "Data not found", nil))
	}

	u.Address = dto.Address
	u.Name = dto.Name
	idb.DB.Save(&u)
	return c.JSON(http.StatusOK, handler.SuccessResponse(u))
}

func (idb *InDB) GetUser(c echo.Context) (err error) {
	id := c.Param("id")

	var u structs.User
	errs := idb.DB.First(&u, id).Error
	if errs != nil {
		return c.JSON(http.StatusNotFound, handler.ErrorResponse(http.StatusNotFound, "Data not found", nil))
	}

	return c.JSON(http.StatusOK, handler.SuccessResponse(u))
}

func (idb *InDB) GetAllUser(c echo.Context) (err error) {
	var users []structs.User

	idb.DB.Find(&users)

	return c.JSON(http.StatusOK, handler.SuccessResponse(users))
}

func (idb *InDB) DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")

	// Check user on database
	var u structs.User
	errs := idb.DB.First(&u, id).Error
	if errs != nil {
		return c.JSON(http.StatusNotFound, handler.ErrorResponse(http.StatusNotFound, "Data not found", nil))
	}

	idb.DB.Delete(&u)
	return c.JSON(http.StatusOK, handler.SuccessResponse("Data has been deleted"))
}
