package config

import (
	"fmt"
	"os"

	"github.com/MatheusMikio/schemas"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initPostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("PostgreSQL")
	if err := godotenv.Load(); err != nil {
		logger.Errorf("Error loading .env file: %v", err)
	}

	dbConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})

	if err != nil {
		logger.Errorf("Db opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(
		&schemas.Client{},
		&schemas.Company{},
		&schemas.Professional{},
		&schemas.Scheduling{},
		&schemas.Service{},
		&schemas.MagicLink{},
	)

	if err != nil {
		logger.Errorf("autoMigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
