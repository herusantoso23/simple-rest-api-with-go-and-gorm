package handler

import (
	"fmt"
	"net/http"
	"simple-rest-api-using-gorm-and-echo/structs"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var message string
	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				message = fmt.Sprintf("%s is not valid email", err.Field())
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
			}
			break
		}
	}
	c.Logger().Error(report)
	c.JSON(report.Code, ErrorResponse(report.Code, message, nil))
}

func SuccessResponse(o interface{}) structs.Result {
	var result structs.Result = structs.Result{
		Code:    http.StatusOK,
		Message: "OK",
		Result:  o,
	}
	return result
}

func ErrorResponse(code int, message string, o interface{}) structs.Result {
	var result structs.Result = structs.Result{
		Code:    code,
		Message: message,
		Result:  o,
	}
	return result
}
