package models

type Service struct {
	ServiceID      int     `json:"serviceID"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Price          float64 `json:"price"`
	Duration       int     `json:"duration"`
	ProfessionalID int     `json:"professionalID"`
}
