package database

import (
	"blockchain/internal/config"
	"blockchain/internal/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (*gorm.DB, error) {
	logger := util.NewLogManager()
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("failed to load config: " + err.Error())
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database: " + err.Error())
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Block{}, &Transaction{})
}
