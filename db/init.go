package db

import (
	"fmt"
	"log"
	"oc-2023/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Psql *gorm.DB
)

func InitDB() (err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DbHost,
		config.DbUser,
		config.DbPass,
		config.DbName,
		config.DbPort,
	)
	Psql, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err = Psql.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	return nil
}
