package models

type Professional struct {
	ProfessionalID int    `json:"professionalID"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	Password       string `json:"password"`
}
