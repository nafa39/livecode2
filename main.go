package main

import (
	"fmt"
	"livecode2/config"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	fmt.Println("Hello World")
}
