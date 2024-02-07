package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	_ "fmt"
)

func NewRequestStore(db *sql.DB) *RequestStore {
	return &RequestStore{db}
}

type RequestStore struct {
	*sql.DB
}

func (rs *RequestStore) RequestAddEstablishment(request models.Request) (int, error) {
	res, err := rs.Exec("INSERT INTO request (professionalID, salon_name, address, city, postal_code, request_status) VALUES (?, ?, ?, ?, ?, 'pending')",
		request.ProfessionalID, request.SalonName, request.Address, request.City, request.PostalCode)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
