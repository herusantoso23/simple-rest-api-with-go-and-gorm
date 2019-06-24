package main

import (
	"simple-rest-api-using-gorm-and-echo/handler"
	"simple-rest-api-using-gorm-and-echo/services"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	e.POST("/user", services.CreateUser)
	e.PUT("/user", services.UpdateUser)
	e.GET("/user", services.GetAllUser)
	e.GET("/user/:id", services.GetUser)
	e.HTTPErrorHandler = handler.ErrorHandler
	e.Logger.Fatal(e.Start(":50000"))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
