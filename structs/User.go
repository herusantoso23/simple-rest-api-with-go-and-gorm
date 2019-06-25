package structs

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}
