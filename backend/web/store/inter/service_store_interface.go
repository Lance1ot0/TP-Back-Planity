package inter

import (
	"TP-Back-Planity/web/models"
)

type ServiceStoreInterface interface {
	AddService(item models.Service) (int, error)
}
