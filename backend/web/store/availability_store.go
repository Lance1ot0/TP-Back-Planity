package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	"fmt"
)

func NewAvailabilityStore(db *sql.DB) *AvailabilityStore {
	return &AvailabilityStore{db}
}

type AvailabilityStore struct {
	*sql.DB
}

func (as *AvailabilityStore) AddAvailability(id int, item models.Availability) (models.Availability, error) {

	existing, err := as.Query("SELECT availabilityID FROM availability WHERE employeeID = ? AND day_of_week = ?",
		id, item.DayOfWeek)
	if err != nil {
		return models.Availability{}, err
	}

	if existing.Next() {
		return models.Availability{}, fmt.Errorf("availability already exists")
	}

	res, err := as.Exec("INSERT INTO availability (employeeID, day_of_week, start_time, end_time) VALUES (?, ?, ?, ?)",
		id, item.DayOfWeek, item.StartTime, item.EndTime)
	if err != nil {
		return models.Availability{}, err
	}

	availabilityID, err := res.LastInsertId()
	if err != nil {
		return models.Availability{}, err
	}

	item.AvailabilityID = int(availabilityID)
	item.EmployeeID = id
	return item, nil
}
