package internal

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email         string `gorm:"unique"`
	PasswordHash  string
	LastLoginDate string
	JWTToken      string
}

// GetUserByEmail retrieves a user by their email

func GetUserByEmail(email string) (*User, error) {

	// This is a mock implementation. Replace with actual database query.

	if email == "test@example.com" {

		return &User{

			Email: "test@example.com",

			PasswordHash: "$2a$10$7a8b9c8d7e6f5g4h3i2j1k0l9m8n7o6p5q4r3s2t1u0v9w8x7y6z5", // bcrypt hash for "password"

		}, nil

	}

	return nil, errors.New("user not found")

}
