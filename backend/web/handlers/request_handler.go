package handlers

import (
	"encoding/json"
	"net/http"

	"TP-Back-Planity/web/models"
	// "TP-Back-Planity/web/utils"
)

func (h *Handler) RequestAddEstablishment() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		item := models.Request{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := h.Store.Request.RequestAddEstablishment(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status    string `json:"status"`
			RequestID int    `json:"requestID"`
		}{
			Status:    "success",
			RequestID: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
