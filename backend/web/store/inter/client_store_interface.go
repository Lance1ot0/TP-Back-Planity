package inter

import (
	"TP-Back-Planity/web/models"
)

type ClientStoreInterface interface {
	GetClient() ([]models.Client, error)
	GetClientById(id int) (models.Client, error)
	GetClientByEmail(email string) (models.Client, error)
	AddClient(client models.Client) (int, error)
	GetPasswordHash(id int) (string, error)
}
