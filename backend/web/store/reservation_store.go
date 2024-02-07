package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	_ "fmt"
)

func NewReservationStore(db *sql.DB) *ReservationStore {
	return &ReservationStore{db}
}

type ReservationStore struct {
	*sql.DB
}

func (rs *ReservationStore) GetAllReservations(id int) ([]models.Reservation, error) {
	rows, err := rs.Query("SELECT * FROM reservation WHERE hairSalonID = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation
	for rows.Next() {
		var reservation models.Reservation
		err := rows.Scan(&reservation.ReservationID, &reservation.TimeSlotID, &reservation.ClientID, &reservation.ServiceID, &reservation.HairSalonID, &reservation.ReservationDate, &reservation.ReservationStatus)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}