package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
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
	result, err := es.Exec("INSERT INTO employee (firstname, lastname, hairSalonID) VALUES (?, ?, ?)",
		employee.Firstname, employee.Lastname, employee.HairSalonID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (es *EmployeeStore) GetAllEmployee(id int) ([]models.Employee, error) {
	var employees []models.Employee

	rows, err := es.Query("SELECT * FROM employee WHERE hairSalonID = ?", id)
	if err != nil {
		return []models.Employee{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		if err = rows.Scan(&employee.EmployeeID, &employee.Firstname, &employee.Lastname, &employee.HairSalonID); err != nil {
			return []models.Employee{}, err
		}
		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return []models.Employee{}, err
	}

	return employees, nil
}
