package schemas

import "gorm.io/gorm"

type CardData struct {
	gorm.Model
	CardBrand    string `json:"cardBrand"`
	CardLast4    string `json:"cardLast4"`
	CardExpMonth int    `json:"cardExpMonth"`
	CardExpYear  int    `json:"cardExpYear"`
}
