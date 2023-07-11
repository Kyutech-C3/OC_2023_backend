package config

import (
	"fmt"
	"log"
	"oc-2023/pkg/domain/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbHost  string
	dbUser  string
	dbPass  string
	dbName  string
	dbPort  string
	APIHost string
)

var (
	db  *gorm.DB
	err error
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&model.Likes{}, &model.Comments{}); err != nil {
		log.Fatal(err)
	}

	return db
}

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("failed load env")
	}
	dbHost = os.Getenv("POSTGRES_HOST")
	dbUser = os.Getenv("POSTGRES_USER")
	dbPass = os.Getenv("POSTGRES_PASSWORD")
	dbName = os.Getenv("POSTGRES_DB")
	dbPort = os.Getenv("POSTGRES_PORT")
	APIHost = os.Getenv("API_HOST")
}
