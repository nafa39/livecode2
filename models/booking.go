package internal

import "github.com/jinzhu/gorm"

type Booking struct {
	gorm.Model
	CustomerID    uint `gorm:"foreignkey:CustomerID"`
	BookingDate   string
	BookingStatus string
}
