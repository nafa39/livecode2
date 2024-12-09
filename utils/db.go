//package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func InitDB() *gorm.DB {
// 	if DB != nil {
// 		return DB
// 	}
// 	var err error
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
// 		PrepareStmt: false,
// 	})
// 	if err != nil {
// 		log.Print("Error connecting to database")
// 		panic(err)
// 	}

// 	log.Print("Database connected")

// 	return DB
// }

// func CloseDB() {
// 	db, err := DB.DB()
// 	if err != nil {
// 		log.Print("Error closing database")
// 		panic(err)
// 	}

// 	db.Close()
// }

package utils

import (
	"fmt"
	internal "livecode2/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&internal.User{}, &internal.Customer{}, &internal.Tour{}, &internal.Booking{}, &internal.TourBooking{}, &internal.Payment{})
	DB = db
}
