package schemas

import (
	base "github.com/MatheusMikio/Base"
	"github.com/MatheusMikio/models"
)

type Client struct {
	base.BaseUser
	models.CardData `json:"cardData"`
	Appointments    []Scheduling `json:"appointments"`
}
