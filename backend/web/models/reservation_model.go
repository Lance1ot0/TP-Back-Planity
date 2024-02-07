package models

type Reservation struct {
	ReservationID     int    `json:"reservationID"`
	TimeSlotID        int    `json:"timeSlotID"`
	ClientID          int    `json:"clientID"`
	ServiceID         int    `json:"serviceID"`
	HairSalonID       int    `json:"hairSalonID"`
	ReservationDate   string `json:"reservationDate"`
	ReservationStatus string `json:"reservationStatus"`
}
