package inter

import (
	"TP-Back-Planity/web/models"
)

type ProfessionalStoreInterface interface {
	GetProfessional() ([]models.Professional, error)
	GetProfessionalById(id int) (models.Professional, error)
	GetProfessionalByEmail(email string) (models.Professional, error)
	AddProfessional(professional models.Professional) (bool, error)
	GetPasswordHash(id int) (string, error)
	GetHairSalon(id int) (models.HairSalon, error)
	GetRequest(id int) (models.Request, error)
	GetEmployeeAvailability(id int) ([]models.Availability, error)
	AddEmployeeAvailability(id int, availability models.Availability) (bool, error)
	GetHairSalonService(id int) ([]models.Service, error)
	GetHairSalonReservation(id int) ([]models.ReservationWithNames, error)
	ListReservationsForPro(hairSalonId int) ([]models.Reservation, error)
}
