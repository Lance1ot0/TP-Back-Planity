package store

import (
	"TP-Back-Planity/web/models"
	"database/sql"
)

func NewAdminStore(db *sql.DB) *AdminStore {
	return &AdminStore{db}
}

type AdminStore struct {
	*sql.DB
}

func (as *AdminStore) GetAdmin() ([]models.Administrator, error) {
	var administrators []models.Administrator

	rows, err := as.Query("SELECT administratorID, firstname, lastname, email, password FROM administrator WHERE administratorID")
	if err != nil {
		return []models.Administrator{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var administrator models.Administrator
		if err = rows.Scan(&administrator.AdministratorID, &administrator.Firstname, &administrator.Lastname, &administrator.Email, &administrator.Password); err != nil {
			return []models.Administrator{}, err
		}
		administrators = append(administrators, administrator)
	}

	if err = rows.Err(); err != nil {
		return []models.Administrator{}, err
	}

	return administrators, nil
}

func (as *AdminStore) GetAdminById(id int) (models.Administrator, error) {
	var administrator models.Administrator

	rows, err := as.Query("SELECT administratorID, firstname, lastname, email, password FROM administrator WHERE administratorID = ?", id)
	if err != nil {
		return models.Administrator{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&administrator.AdministratorID, &administrator.Firstname, &administrator.Lastname, &administrator.Email); err != nil {
			return models.Administrator{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return models.Administrator{}, err
	}

	return administrator, nil
}

func (as *AdminStore) GetAdminByEmail(email string) (models.Administrator, error) {
	var administrator models.Administrator

	rows, err := as.Query("SELECT administratorID, firstname, lastname, email, password FROM administrator WHERE email = ?", email)
	if err != nil {
		return models.Administrator{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&administrator.AdministratorID, &administrator.Firstname, &administrator.Lastname, &administrator.Email, &administrator.Password); err != nil {
			return models.Administrator{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return models.Administrator{}, err
	}

	return administrator, nil
}

func (as *AdminStore) LoginAdmin(email string, password string) (int, error) {
	var id int
	err := as.QueryRow("SELECT administratorID FROM administrator WHERE email = ? AND password = ?", email, password).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}