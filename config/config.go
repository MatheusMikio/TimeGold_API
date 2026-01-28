package config

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = initPostgreSQL()
	if err != nil {
		return fmt.Errorf("error initializing db: %v", err)
	}
	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}

func GetStripeKey() string {
	return os.Getenv("STRIPE_SECRET_KEY")
}

func GetGoogleClientId() string {
	return os.Getenv("GOOGLE_CLIENT_ID")
}

func GetGoogleClientSecret() string {
	return os.Getenv("GOOGLE_CLIENT_SECRET")
}

func GetGoogleRedirectURL(entityType string) string {
	switch entityType {
	case "client":
		return os.Getenv("GOOGLE_REDIRECT_URL_CLIENT")
	case "professional":
		return os.Getenv("GOOGLE_REDIRECT_URL_PROFESSIONAL")
	default:
		return ""
	}
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}

func GetJwtExpirationHours() int {
	hours, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_HOURS"))
	if err != nil {
		return 2
	}
	return hours
}
