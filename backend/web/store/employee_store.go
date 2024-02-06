package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	"fmt"
)

func NewEmployeeStore(db *sql.DB) *EmployeeStore {
	return &EmployeeStore{db}
}

type EmployeeStore struct {
	*sql.DB
}

func (es *EmployeeStore) employeeExists(professionalID int) (bool, error) {
	row := es.QueryRow("SELECT COUNT(*) FROM employee WHERE professionalID = ?", professionalID)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (es *EmployeeStore) AddEmployee(employee models.Employee) (int, error) {
	if exists, err := es.employeeExists(employee.ProfessionalID); err != nil {
		return 0, err
	} else if exists {
		return 0, fmt.Errorf("employee already exists")
	}

	result, err := es.Exec("INSERT INTO employee (firstname, lastname, professionalID, hairSalonID) VALUES (?, ?, ?, ?)",
		employee.Firstname, employee.Lastname, employee.ProfessionalID, employee.HairSalonID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
