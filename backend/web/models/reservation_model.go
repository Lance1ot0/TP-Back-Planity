package models

type Reservation struct {
	ReservationID     int    `json:"reservationID"`
	TimeSlotID        int    `json:"timeSlotID"`
	ClientID          int    `json:"clientID"`
	ServiceID         int    `json:"serviceID"`
	ReservationDate   string `json:"reservationDate"`
	ReservationStatus string `json:"reservationStatus"`
}
