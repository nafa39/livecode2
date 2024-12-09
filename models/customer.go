package internal

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	UserID      uint `gorm:"foreignkey:UserID"`
	Name        string
	Email       string
	PhoneNumber string
	Address     string
}
