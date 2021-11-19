package configs

import (
	"os"
)

type DbConfig struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
}

func NewDBConfig(debug bool) *DbConfig {
	db := DbConfig{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	if debug {
		db.DB_HOST = "localhost"
		db.DB_USER = "joungsik"
		db.DB_PASSWORD = "wjdtlr21"
	}

	return &db
}
