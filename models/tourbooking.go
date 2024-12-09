package internal

import "github.com/jinzhu/gorm"

type TourBooking struct {
	gorm.Model
	BookingID      uint `gorm:"foreignkey:BookingID"`
	TourID         uint `gorm:"foreignkey:TourID"`
	TourStartDate  string
	TourEndDate    string
	NumberOfPerson int
}
