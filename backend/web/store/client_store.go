package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
)

func NewClientStore(db *sql.DB) *ClientStore {
	return &ClientStore{db}
}

type ClientStore struct {
	*sql.DB
}

func (cs *ClientStore) GetClient() ([]models.Client, error) {
	var clients []models.Client

	rows, err := cs.Query("SELECT clientID, firstname, lastname, email FROM client WHERE clientID")
	if err != nil {
		return []models.Client{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var client models.Client
		if err = rows.Scan(&client.ClientID, &client.Firstname, &client.Lastname, &client.Email); err != nil {
			return []models.Client{}, err
		}
		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return []models.Client{}, err
	}

	return clients, nil
}

func (cs *ClientStore) GetClientById(id int) (models.Client, error) {
	var client models.Client

	rows, err := cs.Query("SELECT clientID, firstname, lastname, email FROM client WHERE clientID = ?", id)
	if err != nil {
		return models.Client{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&client.ClientID, &client.Firstname, &client.Lastname, &client.Email); err != nil {
			return models.Client{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return models.Client{}, err
	}

	return client, nil
}
