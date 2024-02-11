package inter

import (
	"TP-Back-Planity/web/models"
)

type ClientStoreInterface interface {
	GetClient() ([]models.Client, error)
	GetClientById(id int) (models.Client, error)
	ResearchHairSalon(name string) ([]models.HairSalon, error)
	GetClientByEmail(email string) (models.Client, error)
	AddClient(client models.Client) (int, error)
	GetPasswordHash(id int) (string, error)
	AddReservation(reservation models.Reservation) (models.Reservation, error)
	ListReservations(clientId int) ([]models.Reservation, error)
	CancelReservation(clientId int) (bool, error)
}
