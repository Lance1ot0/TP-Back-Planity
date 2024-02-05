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

func (as *AdminStore) ListRequests() ([]models.Request, error) {
	var requests []models.Request
	rows, err := as.Query("SELECT * FROM request WHERE request_status = 'pending'")


    if err != nil {
        return nil, err
    }
	
    defer rows.Close()

    for rows.Next() {
        var request models.Request
        if err := rows.Scan(&request.RequestID, &request.ProfessionalID, &request.SalonName, &request.Address, &request.City, &request.PostalCode, &request.RequestDate, &request.RequestStatus); err != nil {
            return nil, err
        }
        requests = append(requests, request)
    }

	
	if err != nil {
		return []models.Request{}, err
	}

	return requests, nil
}

func (as *AdminStore) UpdateRequest(id int, status string) (bool, error) {

	result, err:= as.Exec("UPDATE request SET request_status = ? WHERE RequestID = ?", status, id)

    if err != nil {
        return false, err
    }

	rowsAffected, err := result.RowsAffected()
    if err != nil {
        return false, err
    }

	if rowsAffected == 0 {
        return false, nil
    }

	return true, nil
}