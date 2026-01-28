package services

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/MatheusMikio/config"
	"github.com/MatheusMikio/schemas"
	"github.com/golang-jwt/jwt/v5"
)

func generateToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func generateJWT(userId uint, email, entityType, role string) (string, int64, error) {
	expirationHours := config.GetJwtExpirationHours()
	expiresAt := time.Now().Add(time.Duration(expirationHours) * time.Hour)
	expiresIn := int64(expirationHours * 3600)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     userId,
		"email":       email,
		"entity_type": entityType,
		"role":        role,
		"exp":         expiresAt.Unix(),
		"iat":         time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetJwtSecret()))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresIn, nil
}

func generateJWTProfessional(userID uint, email string, role schemas.ProfessionalRole) (string, int64, error) {
	expirationHours := config.GetJwtExpirationHours()
	expiresAt := time.Now().Add(time.Duration(expirationHours) * time.Hour)
	expiresIn := int64(expirationHours * 3600)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":     userID,
		"email":       email,
		"entity_type": "professional",
		"role":        string(role),
		"exp":         expiresAt.Unix(),
		"iat":         time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetJwtSecret()))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiresIn, nil
}
