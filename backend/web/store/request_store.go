package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	"fmt"
)

func NewRequestStore(db *sql.DB) *RequestStore {
	return &RequestStore{db}
}

type RequestStore struct {
	*sql.DB
}

func (rs *RequestStore) requestExists(professionalID int) (bool, error) {
	row := rs.QueryRow("SELECT COUNT(*) FROM request WHERE professionalID = ?", professionalID)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (rs *RequestStore) RequestAddEstablishment(request models.Request) (models.Request, error) {
	if exists, err := rs.requestExists(request.ProfessionalID); err != nil {
		return models.Request{}, err
	} else if exists {
		return models.Request{}, fmt.Errorf("request already exists")
	}

	res, err := rs.Exec("INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_status) VALUES (?, ?, ?, ?, ?, 'pending')",
		request.ProfessionalID, request.SalonName, request.Address, request.City, request.PostalCode)
	if err != nil {
		return models.Request{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.Request{}, err
	}

	request.RequestID = int(id)
	return request, nil
}
