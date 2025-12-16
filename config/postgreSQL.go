package config

import (
	"github.com/MatheusMikio/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbConnection = "host=localhost port=5432 user=postgres password=Mikio123 dbname=TimeGold"
)

func initPostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("PostgreSQL")
	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})

	if err != nil {
		logger.Errorf("Db opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Client{}, &schemas.Company{}, &schemas.Professional{}, &schemas.Scheduling{}, &schemas.Service{})
	if err != nil {
		logger.Errorf("autoMigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
