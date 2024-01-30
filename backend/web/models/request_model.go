package models

type Request struct {
	RequestID      int    `json:"requestID"`
	ProfessionalID int    `json:"professionalID"`
	SalonName      string `json:"salonName"`
	Address        string `json:"address"`
	City           string `json:"city"`
	PostalCode     string `json:"postalCode"`
	RequestDate    string `json:"requestDate"`
	RequestStatus  string `json:"requestStatus"`
}
