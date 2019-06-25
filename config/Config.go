package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"simple-rest-api-using-gorm-and-echo/structs"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@/go?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.User{})
	return db
}
