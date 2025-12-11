package models

type Adress struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	City       string `json:"city"`
	State      string `json:"state"`
	Cep        string `json:"cep"`
	Country    string `json:"country"`
	Complement string `json:"complement"`
}
