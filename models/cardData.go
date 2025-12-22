package models

import "gorm.io/gorm"

type CardData struct {
	gorm.Model
	StripeCardId string `json:"stripeCardId"`
	CardBrand    string `json:"cardBrand"`
	CardLast4    string `json:"cardLast4"`
	CardExpMonth int    `json:"cardExpMonth"`
	CardExpYear  int    `json:"cardExpYear"`
}
