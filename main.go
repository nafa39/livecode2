package main

import (
	"fmt"
	"livecode2/config"
	handler "livecode2/internal/userhandler"

	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	defer config.CloseDB()
	fmt.Println("Hello World")

	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})

	e.POST("/register", handler.RegisterUser)
	e.Logger.Fatal(e.Start(":8080"))
}
