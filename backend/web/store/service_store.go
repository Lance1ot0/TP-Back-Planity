package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	_ "fmt"
)

func NewServiceStore(db *sql.DB) *ServiceStore {
	return &ServiceStore{db}
}

type ServiceStore struct {
	*sql.DB
}

func (ss *ServiceStore) AddService(item models.Service) (int, error) {
	result, err := ss.Exec("INSERT INTO service (name, description, price, duration, hairSalonID) VALUES (?, ?, ?, ?, ?)",
		item.Name, item.Description, item.Price, item.Duration, item.HairSalonID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (ss *ServiceStore) ListServices(hairSalonId int) ([]models.Service, error) {

	var services []models.Service
	var service models.Service

	rows, err := ss.Query("SELECT * FROM service WHERE hairSalonID = ?", hairSalonId)
	if err != nil {
		return []models.Service{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&service.ServiceID, &service.Name, &service.Description, &service.Price, &service.Duration, &service.HairSalonID); err != nil {
			return []models.Service{}, err
		}

		services = append(services, service)
	}

	if err = rows.Err(); err != nil {
		return []models.Service{}, err
	}

	return services, nil
}
