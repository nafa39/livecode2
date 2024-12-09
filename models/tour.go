package internal

import "github.com/jinzhu/gorm"

type Tour struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Duration    string
}
