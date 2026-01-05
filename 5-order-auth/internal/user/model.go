package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone     string `gorm:"type:varchar(20);unique;not null"`
	Name      string
	SessionId string
	Code      int
}
