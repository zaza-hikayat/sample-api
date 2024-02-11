package lib

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionDB(conf *Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.DbHost,
		conf.DbUsername,
		conf.DbPassword,
		conf.DbName,
		conf.DbPort,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))

	if err != nil {
		log.Fatalf("Error connect database: ", err)
	}
	return db
}
