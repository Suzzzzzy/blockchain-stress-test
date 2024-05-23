package main

import (
	"blockchain/internal/database"
	"blockchain/internal/util"
)

func main() {
	logger := util.NewLogManager()

	db, err := database.Connect()
	if err != nil {
		logger.Error("failed to connect to database: " + err.Error())
	}

	err = database.AutoMigrate(db)
	if err != nil {
		logger.Error("failed to migrate database: " + err.Error())
	} else {
		logger.Info("database migration completed successfully")
	}

}
