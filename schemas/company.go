package schemas

import (
	"github.com/MatheusMikio/models"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	Name               string             `json:"name"`
	Category           string             `json:"category"`
	Email              string             `json:"email"`
	Cnpj               string             `json:"cnpj"`
	Description        string             `json:"description"`
	Phone              string             `json:"phone"`
	MonthlyFee         float64            `json:"monthlyFee"`
	AppointmentFeePct  float64            `json:"appointmentFeePct"`
	LogoURL            string             `json:"logoUrl"`
	BannerURL          string             `json:"bannerUrl"`
	Professionals      []Professional     `json:"professionals" gorm:"foreignKey:CompanyID"`
	Adress             models.Adress      `json:"adress" gorm:"embedded"`
	LegallyResponsible LegallyResponsible `json:"legallyResponsible" gorm:"embedded"`
	Pix                models.Pix         `json:"pix" gorm:"embedded"`
	WorkingHours       WorkingHours       `json:"workingHours" gorm:"type:json"`
	Services           []Service          `json:"services" gorm:"foreignKey:CompanyID"`
	CardDataID         uint               `json:"cardDataId"`
	CardData           models.CardData    `json:"-" gorm:"foreignKey:CardDataID"`
	StripeIDs          StripeIDs          `json:"stripeIds" gorm:"embedded"`
	Appointments       []Scheduling       `json:"appointments"`
}

type LegallyResponsible struct {
	FullName string `json:"fullName"`
	Cpf      string `json:"cpf"`
}

type WorkingHours struct {
	Monday    DaySchedule `json:"monday"`
	Tuesday   DaySchedule `json:"tuesday"`
	Wednesday DaySchedule `json:"wednesday"`
	Thursday  DaySchedule `json:"thursday"`
	Friday    DaySchedule `json:"friday"`
	Saturday  DaySchedule `json:"saturday"`
	Sunday    DaySchedule `json:"sunday"`
}

type DaySchedule struct {
	IsOpen    bool   `json:"isOpen"`
	OpenTime  string `json:"openTime"`
	CloseTime string `json:"closeTime"`
}

type StripeIDs struct {
	StripeCustomerID      string `json:"stripeCustomerId"`
	StripePaymentMethodId string `json:"stripePaymentMethodId"`
	SubscriptionID        string `json:"subscriptionId"`
}
