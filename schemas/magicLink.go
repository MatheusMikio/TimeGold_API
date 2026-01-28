package schemas

import (
	"time"

	"gorm.io/gorm"
)

type MagicLink struct {
	gorm.Model
	Email      string    `json:"email" gorm:"index"`
	Token      string    `json:"token" gorm:"uniqueIndex"`
	EntityType string    `json:"entityType"`
	ExpiresAt  time.Time `json:"expiresAt"`
	Used       bool      `json:"used" gorm:"default:false"`
}
