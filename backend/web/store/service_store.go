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
