package inter

import (
	"TP-Back-Planity/web/models"
)

type ProfessionalStoreInterface interface {
	GetProfessional() ([]models.Professional, error)
	GetProfessionalById(id int) (models.Professional, error)
	GetProfessionalByEmail(email string) (models.Professional, error)
	AddProfessional(professional models.Professional) (int, error)
	GetPasswordHash(id int) (string, error)
}
