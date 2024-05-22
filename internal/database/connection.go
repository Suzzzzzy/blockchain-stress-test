package database

import (
	"blockchain/internal/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Connect() (*gorm.DB, error) {
	logger := util.NewLogManager()

	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect to database: " + err.Error())
	}
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Block{}, &Transaction{})
}
