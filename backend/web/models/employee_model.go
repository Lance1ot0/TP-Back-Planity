package models

type Employee struct {
	EmployeeID  int    `json:"employeeID"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	HairSalonID int    `json:"hairSalonID"`
}

type EmployeeInfo struct {
	EmployeeID           int    `json:"employeeID"`
	Firstname            string `json:"firstname"`
	Lastname             string `json:"lastname"`
	HairSalonID          int    `json:"hairSalonID"`
	Availabilities       []Availability
	EmployeeReservations []Reservation
}
