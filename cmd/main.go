package main

import (
	"blockcain/internal/config"
	"blockcain/internal/database"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("failed to connect to database: %v", err)
	}

	err = database.AutoMigrate(db)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

}

