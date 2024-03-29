package inter

import (
	"TP-Back-Planity/web/models"
)

type AdminStoreInterface interface {
	GetAdmin() ([]models.Administrator, error)
	GetAdminById(id int) (models.Administrator, error)
	GetAdminByEmail(email string) (models.Administrator, error)
	LoginAdmin(username, password string) (int, error)
	ListRequests() ([]models.Request, error)
	UpdateRequest(id int, status string) (bool, error)
	GetRequestById(id int) (models.Request, error)
	CreateSalon(models.Request) (int, error)
}
