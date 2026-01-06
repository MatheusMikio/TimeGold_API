package company

import "github.com/MatheusMikio/models"

type CompanyRequest struct {
	Name               string                    `json:"name"`
	Category           string                    `json:"category"`
	Email              string                    `json:"email"`
	Cnpj               string                    `json:"cnpj"`
	Description        string                    `json:"description"`
	Phone              string                    `json:"phone"`
	MonthlyFee         float64                   `json:"monthlyFee"`
	AppointmentFeePct  float64                   `json:"appointmentFeePct"`
	LogoURL            string                    `json:"logoUrl"`
	BannerURL          string                    `json:"bannerUrl"`
	Adress             models.Adress             `json:"adress"`
	LegallyResponsible LegallyResponsibleRequest `json:"legallyResponsible"`
	Pix                models.Pix                `json:"pix"`
	WorkingHours       WorkingHoursRequest       `json:"workingHours"`
	CardDataID         uint                      `json:"cardDataId"`
	StripeIds          StripeIDsRequest          `json:"stripeIds"`
}

type LegallyResponsibleRequest struct {
	FullName string `json:"fullName"`
	Cpf      string `json:"cpf"`
}

type DayScheduleRequest struct {
	IsOpen    bool   `json:"isOpen"`
	OpenTime  string `json:"openTime"`
	CloseTime string `json:"closeTime"`
}

type WorkingHoursRequest struct {
	Monday    DayScheduleRequest `json:"monday"`
	Tuesday   DayScheduleRequest `json:"tuesday"`
	Wednesday DayScheduleRequest `json:"wednesday"`
	Thursday  DayScheduleRequest `json:"thursday"`
	Friday    DayScheduleRequest `json:"friday"`
	Saturday  DayScheduleRequest `json:"saturday"`
	Sunday    DayScheduleRequest `json:"sunday"`
}

type StripeIDsRequest struct {
	StripeCustomerID      string `json:"stripeCustomerId"`
	StripePaymentMethodId string `json:"stripePaymentMethodId"`
	SubscriptionID        string `json:"subscriptionId"`
}

type UpdateCompanyRequest struct {
	ID                 uint                      `json:"id"`
	Name               string                    `json:"name"`
	Category           string                    `json:"category"`
	Email              string                    `json:"email"`
	Cnpj               string                    `json:"cnpj"`
	Description        string                    `json:"description"`
	Phone              string                    `json:"phone"`
	MonthlyFee         float64                   `json:"monthlyFee"`
	AppointmentFeePct  float64                   `json:"appointmentFeePct"`
	LogoURL            string                    `json:"logoUrl"`
	BannerURL          string                    `json:"bannerUrl"`
	Adress             models.Adress             `json:"adress"`
	LegallyResponsible LegallyResponsibleRequest `json:"legallyResponsible"`
	Pix                models.Pix                `json:"pix"`
	WorkingHours       WorkingHoursRequest       `json:"workingHours"`
	CardDataID         uint                      `json:"cardDataId"`
	StripeIds          StripeIDsRequest          `json:"stripeIds"`
}
