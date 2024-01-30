package models

type Employee struct {
	EmployeeID     int    `json:"employeeID"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	ProfessionalID int    `json:"professionalID"`
	HairSalonID    int    `json:"hairSalonID"`
}
