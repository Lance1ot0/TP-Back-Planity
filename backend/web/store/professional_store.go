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

func (ps *ProfessionalStore) AddProfessional(professional models.Professional) (int, error) {
	if professional.Firstname == "" || professional.Lastname == "" || professional.Email == "" ||
		professional.Phone == "" || professional.Address == "" || professional.Password == "" {
		return 0, fmt.Errorf("All fields must be completed")
	}

	pro, _ := ps.GetProfessionalByEmail(professional.Email)
	if pro.Email != "" {
		return 0, fmt.Errorf("Email already exist")
	}

	res, err := ps.Exec("INSERT INTO professional (firstname, lastname, email, phone, address, password) VALUES (?, ?, ?, ?, ?, ?)", professional.Firstname, professional.Lastname, professional.Email, professional.Phone, professional.Address, professional.Password)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
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
