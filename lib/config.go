package lib

import (
	"os"

	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

type Config struct {
	DbHost     string
	DbPassword string
	DbUsername string
	DbName     string
	DbPort     string
	AppPort    string
}

func LoadConfig() *Config {
	return &Config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbName:     os.Getenv("DB_NAME"),
		DbPort:     os.Getenv("DB_PORT"),
		AppPort:    os.Getenv("APP_PORT"),
	}
}
