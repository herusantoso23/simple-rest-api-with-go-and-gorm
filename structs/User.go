package structs

type User struct {
	ID      int    `json:"id" gorm:"AUTO_INCREMENT;PRIMARY_KEY;Column:id"`
	Name    string `json:"name" validate:"required" gorm:"Column:name;size:100"`
	Address string `json:"address" validate:"required" gorm:"Column:address;size:200"`
}
