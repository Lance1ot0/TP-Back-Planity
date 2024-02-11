package inter

import (
	"TP-Back-Planity/web/models"
)

type RequestStoreInterface interface {
	RequestAddEstablishment(request models.Request) (models.Request, error)
}
