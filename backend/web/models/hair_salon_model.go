package models

type HairSalon struct {
	HairSalonID    int    `json:"hairSalonID"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	ProfessionalID int    `json:"professionalID"`
}
