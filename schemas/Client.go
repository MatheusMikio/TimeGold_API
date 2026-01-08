package schemas

import (
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/models/base"
)

type Client struct {
	base.BaseUser
	CardDataID   uint             `json:"cardDataId" gorm:"index"`
	CardData     *models.CardData `json:"cardData" gorm:"foreignKey:CardDataID"`
	Appointments []Scheduling     `json:"appointments"`
}
