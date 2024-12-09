package internal

import "github.com/jinzhu/gorm"

type Payment struct {
	gorm.Model
	BookingID     uint `gorm:"foreignkey:BookingID"`
	PaymentDate   string
	PaymentAmount float64
}
