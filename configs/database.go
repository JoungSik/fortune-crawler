package configs

import (
	"crawler/cmd/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(debug bool) (*gorm.DB, error) {
	config := NewDBConfig(debug)

	dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASSWORD + " dbname=KakaoBot port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Fortune{})

	return db, nil
}
