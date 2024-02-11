package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"TP-Back-Planity/web/models"

	"github.com/go-chi/chi"
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

func (h *Handler) GetAllEmployee() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		employees, err := h.Store.Employee.GetAllEmployee(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(employees)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
