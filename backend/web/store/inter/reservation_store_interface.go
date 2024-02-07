package inter

import (
	"TP-Back-Planity/web/models"
)

type ReservationStoreInterface interface {
	GetAllReservations(id int) ([]models.Reservation, error)
}
