package handlers

import (
	"encoding/json"
	"net/http"

	"TP-Back-Planity/web/models"
)

func (h *Handler) AddEmploye() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		var employee models.Employee
		err := json.NewDecoder(request.Body).Decode(&employee)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := h.Store.Employee.AddEmployee(employee)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status     string `json:"status"`
			EmployeeID int    `json:"employeeID"`
		}{
			Status:     "success",
			EmployeeID: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
