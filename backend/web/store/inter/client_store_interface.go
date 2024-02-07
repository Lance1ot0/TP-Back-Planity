package inter

import (
	"TP-Back-Planity/web/models"
)

type ClientStoreInterface interface {
	GetClient() ([]models.Client, error)
	GetClientById(id int) (models.Client, error)
	GetHairSalon(name string) ([]models.HairSalon, error)
}
