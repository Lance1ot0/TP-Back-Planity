package inter

import (
	"TP-Back-Planity/web/models"
)

type AvailabilityStoreInterface interface {
	AddAvailability(id int, item models.Availability) (models.Availability, error)
}
