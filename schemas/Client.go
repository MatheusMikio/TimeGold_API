package schemas

import (
	"github.com/MatheusMikio/models"
	"github.com/MatheusMikio/models/base"
)

type Client struct {
	base.BaseUser
	models.CardData `json:"cardData"`
	Appointments    []Scheduling `json:"appointments"`
}
