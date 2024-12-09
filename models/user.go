package internal

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email         string `gorm:"unique"`
	PasswordHash  string
	LastLoginDate string
	JWTToken      string
}
