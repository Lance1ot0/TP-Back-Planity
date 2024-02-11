package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	"fmt"
)

func NewProfessionalStore(db *sql.DB) *ProfessionalStore {
	return &ProfessionalStore{db}
}

type ProfessionalStore struct {
	*sql.DB
}

func (ps *ProfessionalStore) GetProfessional() ([]models.Professional, error) {
	var professionals []models.Professional

	rows, err := ps.Query("SELECT professionalID, firstname, lastname, email, phone, address FROM professional WHERE professionalID")
	if err != nil {
		return []models.Professional{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var professional models.Professional
		if err = rows.Scan(&professional.ProfessionalID, &professional.Firstname, &professional.Lastname, &professional.Email, &professional.Phone, &professional.Address); err != nil {
			return []models.Professional{}, err
		}
		professionals = append(professionals, professional)
	}

	if err = rows.Err(); err != nil {
		return []models.Professional{}, err
	}

	return professionals, nil
}

func (ps *ProfessionalStore) GetProfessionalById(id int) (models.Professional, error) {
	var professional models.Professional

	rows, err := ps.Query("SELECT professionalID, firstname, lastname, email, phone, address FROM professional WHERE professionalID = ?", id)
	if err != nil {
		return models.Professional{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&professional.ProfessionalID, &professional.Firstname, &professional.Lastname, &professional.Email); err != nil {
			return models.Professional{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return models.Professional{}, err
	}

	return professional, nil
}

func (ps *ProfessionalStore) GetProfessionalByEmail(email string) (models.Professional, error) {
	var professional models.Professional

	rows, err := ps.Query("SELECT professionalID, firstname, lastname, email, phone, address FROM professional WHERE email = ?", email)
	if err != nil {
		return models.Professional{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&professional.ProfessionalID, &professional.Firstname, &professional.Lastname, &professional.Email, &professional.Phone, &professional.Address); err != nil {
			return models.Professional{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return models.Professional{}, err
	}

	return professional, nil
}

func (ps *ProfessionalStore) GetHairSalon(id int) (models.HairSalon, error) {
	var hairSalon models.HairSalon

	rows, err := ps.Query("SELECT hairSalonID, name, address, city, postal_code, professionalID FROM hairSalon WHERE professionalID = ?", id)
	if err != nil {
		return models.HairSalon{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&hairSalon.HairSalonID, &hairSalon.Name, &hairSalon.Address, &hairSalon.City, &hairSalon.PostalCode, &hairSalon.ProfessionalID); err != nil {
			return models.HairSalon{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return models.HairSalon{}, err
	}

	return hairSalon, nil
}

func (ps *ProfessionalStore) GetRequest(id int) (models.Request, error) {
	var request models.Request

	rows, err := ps.Query("SELECT requestID, professionalID, salon_name, address, city, postal_code, request_date, request_status FROM request WHERE professionalID = ?", id)
	if err != nil {
		return models.Request{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&request.RequestID, &request.ProfessionalID, &request.SalonName, &request.Address, &request.City, &request.PostalCode, &request.RequestDate, &request.RequestStatus); err != nil {
			return models.Request{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return models.Request{}, err
	}

	return request, nil
}

func (ps *ProfessionalStore) AddProfessional(professional models.Professional) (bool, error) {
	if professional.Firstname == "" || professional.Lastname == "" || professional.Email == "" ||
		professional.Phone == "" || professional.Address == "" || professional.Password == "" {
		return false, fmt.Errorf("All fields must be completed")
	}

	pro, _ := ps.GetProfessionalByEmail(professional.Email)
	if pro.Email != "" {
		return false, fmt.Errorf("Email already exist")
	}

	_, err := ps.Exec("INSERT INTO professional (firstname, lastname, email, phone, address, password) VALUES (?, ?, ?, ?, ?, ?)", professional.Firstname, professional.Lastname, professional.Email, professional.Phone, professional.Address, professional.Password)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ps *ProfessionalStore) GetPasswordHash(id int) (string, error) {
	var password string

	rows, err := ps.Query("SELECT password FROM professional WHERE professionalID = ?", id)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&password); err != nil {
			return "", err
		}
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return password, nil
}

func (ps *ProfessionalStore) GetEmployeeAvailability(id int) ([]models.Availability, error) {
	var availabilities []models.Availability

	rows, err := ps.Query("SELECT * FROM availability WHERE employeeID = ?", id)
	if err != nil {
		return []models.Availability{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var availability models.Availability
		if err = rows.Scan(&availability.AvailabilityID, &availability.EmployeeID, &availability.DayOfWeek, &availability.StartTime, &availability.EndTime, &availability.IntervalTime); err != nil {
			return []models.Availability{}, err
		}
		availabilities = append(availabilities, availability)
	}

	if err = rows.Err(); err != nil {
		return []models.Availability{}, err
	}

	return availabilities, nil
}

func (ps *ProfessionalStore) AddEmployeeAvailability(id int, availability models.Availability) (bool, error) {
	var existingAvailability models.Availability
	err := ps.QueryRow("SELECT * FROM availability WHERE employeeID = ? AND day_of_week = ?", id, availability.DayOfWeek).Scan(
		&existingAvailability.AvailabilityID,
		&existingAvailability.EmployeeID,
		&existingAvailability.DayOfWeek,
		&existingAvailability.StartTime,
		&existingAvailability.EndTime,
		&existingAvailability.IntervalTime,
	)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err == sql.ErrNoRows {
		_, err = ps.Exec(`
			INSERT INTO availability (employeeID, day_of_week, start_time, end_time, interval_time)
			VALUES (?, ?, ?, ?, ?)
		`, id, availability.DayOfWeek, availability.StartTime, availability.EndTime, availability.IntervalTime)
		if err != nil {
			return false, err
		}
	} else {
		_, err = ps.Exec(`
			UPDATE availability
			SET start_time = ?, end_time = ?
			WHERE availabilityID = ?
		`, availability.StartTime, availability.EndTime, existingAvailability.AvailabilityID)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
