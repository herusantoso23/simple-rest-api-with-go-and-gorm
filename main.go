package main

import (
	"simple-rest-api-using-gorm-and-echo/config"
	"simple-rest-api-using-gorm-and-echo/handler"
	"simple-rest-api-using-gorm-and-echo/services"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	db := config.DBInit()
	inDB := &services.InDB{DB: db}
	e := echo.New()
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	e.POST("/user", inDB.CreateUser)
	e.PUT("/user/:id", inDB.UpdateUser)
	e.GET("/user", inDB.GetAllUser)
	e.GET("/user/:id", inDB.GetUser)
	e.DELETE("/user/:id", inDB.DeleteUser)
	e.HTTPErrorHandler = handler.ErrorHandler
	e.Logger.Fatal(e.Start(":50000"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
