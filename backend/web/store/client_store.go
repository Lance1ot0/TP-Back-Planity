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

func (cs *ClientStore) ResearchHairSalon(name string) ([]models.HairSalon, error) {

	var salons []models.HairSalon

	query := "SELECT * FROM hairSalon WHERE name LIKE ?"

	rows, err := cs.Query(query, name+"%")

	if err != nil {
		return []models.HairSalon{}, err
	}

	var salon models.HairSalon

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&salon.HairSalonID, &salon.Name, &salon.Address, &salon.City, &salon.PostalCode, &salon.ProfessionalID); err != nil {
			return []models.HairSalon{}, err
		}
		salons = append(salons, salon)
	}

	if err = rows.Err(); err != nil {
		return []models.HairSalon{}, err
	}

	return salons, nil
}

func (cs *ClientStore) GetClientByEmail(email string) (models.Client, error) {
	var client models.Client

	rows, err := cs.Query("SELECT clientID, firstname, lastname, email FROM client WHERE email = ?", email)
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

func (cs *ClientStore) AddReservation(reservation models.Reservation) (models.Reservation, error) {

	res, err := cs.Exec("INSERT INTO reservation (employeeID, clientID, serviceID, hairSalonID, date) VALUES (?, ?, ?, ?, ?)", reservation.EmployeeID, reservation.HairSalonID, reservation.ClientID, reservation.ServiceID, reservation.ReservationDate)
	if err != nil {
		return models.Reservation{}, err
	}

	reservationID, err := res.LastInsertId()
	fmt.Println(res.RowsAffected())
	if err != nil {
		return models.Reservation{}, err
	}

	reservation.ReservationID = int(reservationID)
	reservation.ReservationStatus = "confirmed"

	return reservation, nil
}

func (cs *ClientStore) ListReservations(clientId int) ([]models.Reservation, error) {

	var reservations []models.Reservation
	var reservation models.Reservation

	rows, err := cs.Query("SELECT * FROM reservation WHERE clientID = ?", clientId)
	if err != nil {
		return []models.Reservation{}, err
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&reservation.ReservationID, &reservation.EmployeeID, &reservation.HairSalonID, &reservation.ClientID, &reservation.ServiceID, &reservation.ReservationDate, &reservation.ReservationStatus); err != nil {
			return []models.Reservation{}, err
		}

		reservations = append(reservations, reservation)
	}

	if err = rows.Err(); err != nil {
		return []models.Reservation{}, err
	}

	return reservations, nil
}

func (cs *ClientStore) CancelReservation(reservationId int) (bool, error) {

	result, err := cs.Exec("UPDATE reservation SET status = 'canceled' WHERE reservationID = ?", reservationId)

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
func (cs *ClientStore) GetEmployeesWithAvailabilities(hairSalonId int) ([]models.EmployeeInfo, error) {
	var employeesMap = make(map[int]*models.EmployeeInfo)

	rows, err := cs.Query(`
        SELECT e.employeeID, e.firstname, e.lastname, e.hairSalonID,
               a.availabilityID, a.day_of_week, a.start_time, a.end_time,
               r.date, r.hairSalonID, r.serviceID, r.clientID, r.employeeID
        FROM employee e
        LEFT JOIN availability a ON e.employeeID = a.employeeID
        LEFT JOIN reservation r ON r.employeeID = e.employeeID
        WHERE e.hairSalonID = ?
    `, hairSalonId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employeeID int
		var firstname, lastname string
		var hairSalonID int
		var availabilityID, reservationID sql.NullInt64
		var dayOfWeek, startTime, endTime, reservationDateTime sql.NullString
		var reservationHairSalonID, serviceID, clientID sql.NullInt64

		err := rows.Scan(&employeeID, &firstname, &lastname, &hairSalonID,
			&availabilityID, &dayOfWeek, &startTime, &endTime,
			&reservationDateTime, &reservationID, &reservationHairSalonID, &serviceID, &clientID)
		if err != nil {
			return nil, err
		}

		// Créer un nouvel employé s'il n'existe pas encore dans la carte
		if _, ok := employeesMap[employeeID]; !ok {
			employeesMap[employeeID] = &models.EmployeeInfo{
				EmployeeID:           employeeID,
				Firstname:            firstname,
				Lastname:             lastname,
				HairSalonID:          hairSalonID,
				Availabilities:       make([]models.Availability, 0),
				EmployeeReservations: make([]models.Reservation, 0),
			}
		}

		// Ajouter la disponibilité à l'employé correspondant
		if availabilityID.Valid {
			employeesMap[employeeID].Availabilities = append(employeesMap[employeeID].Availabilities, models.Availability{
				AvailabilityID: int(availabilityID.Int64),
				EmployeeID:     employeeID,
				DayOfWeek:      dayOfWeek.String,
				StartTime:      startTime.String,
				EndTime:        endTime.String,
			})
		}

		// Ajouter la réservation à l'employé correspondant
		if reservationID.Valid {
			employeesMap[employeeID].EmployeeReservations = append(employeesMap[employeeID].EmployeeReservations, models.Reservation{
				ReservationID:   int(reservationID.Int64),
				EmployeeID:      employeeID,
				ReservationDate: reservationDateTime.String,
				HairSalonID:     int(reservationHairSalonID.Int64),
				ServiceID:       int(serviceID.Int64),
				ClientID:        int(clientID.Int64),
			})
		}
	}

	// Convertir la carte d'employés en une liste
	var employees []models.EmployeeInfo
	for _, employee := range employeesMap {
		employees = append(employees, *employee)
	}

	return employees, nil
}

func (cs *ClientStore) GetHairSalonById(id int) (models.HairSalon, error) {
	var hairSalon models.HairSalon

	rows, err := cs.Query("SELECT hairSalonID, name, address, city, postal_code, professionalID FROM hairSalon WHERE hairSalonID = ?", id)
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
