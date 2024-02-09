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
