package handlers

import (
	"encoding/json"
	"net/http"

	"TP-Back-Planity/web/models"
)

func (h *Handler) AddService() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		item := models.Service{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := h.Store.Service.AddService(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status    string `json:"status"`
			ServiceID int    `json:"serviceID"`
		}{
			Status:    "success",
			ServiceID: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
