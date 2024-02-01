package models

type Administrator struct {
	AdministratorID int    `json:"administratorID"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
}
