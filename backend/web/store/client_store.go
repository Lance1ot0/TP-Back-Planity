package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
	"fmt"
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

func (ps *ClientStore) GetClientByEmail(email string) (models.Client, error) {
	var client models.Client

	rows, err := ps.Query("SELECT clientID, firstname, lastname, email FROM client WHERE email = ?", email)
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

func (cs *ClientStore) AddClient(client models.Client) (int, error) {
	if client.Firstname == "" || client.Lastname == "" || client.Email == "" || client.Password == "" {
		return 0, fmt.Errorf("All fields must be completed")
	}

	pro, _ := cs.GetClientByEmail(client.Email)
	if pro.Email != "" {
		return 0, fmt.Errorf("Email already exist")
	}

	res, err := cs.Exec("INSERT INTO client (firstname, lastname, email, password) VALUES (?, ?, ?, ?)", client.Firstname, client.Lastname, client.Email, client.Password)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (cs *ClientStore) GetPasswordHash(id int) (string, error) {
	var password string

	rows, err := cs.Query("SELECT password FROM client WHERE clientID = ?", id)
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
