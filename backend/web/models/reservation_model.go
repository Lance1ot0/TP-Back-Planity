package models

type Reservation struct {
	ReservationID     int    `json:"reservationID"`
	EmployeeID        int    `json:"employeeID"`
	HairSalonID       int    `json:"hairSalonID"`
	ClientID          int    `json:"clientID"`
	ServiceID         int    `json:"serviceID"`
	ReservationDate   string `json:"reservationDate"`
	ReservationStatus string `json:"reservationStatus"`
}

type ReservationWithNames struct {
	ReservationID     int    `json:"reservationID"`
	EmployeeID        int    `json:"employeeID"`
	HairSalonID       int    `json:"hairSalonID"`
	ClientID          int    `json:"clientID"`
	ServiceID         int    `json:"serviceID"`
	ReservationDate   string `json:"reservationDate"`
	ReservationStatus string `json:"reservationStatus"`
	EmployeeFirstname string `json:"employeeFirstname"`
	EmployeeLastname  string `json:"employeeLastname"`
	ServiceName       string `json:"serviceName"`
	ClientFirstname   string `json:"clientFirstname"`
	ClientLastname    string `json:"clientLastname"`
}
