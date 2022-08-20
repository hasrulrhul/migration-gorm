package main

import (
	"log"

	"github.com/hasrulrhul/migration-gorm/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()

}
