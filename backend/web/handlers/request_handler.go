package handlers

import (
	"encoding/json"
	"net/http"

	"TP-Back-Planity/web/middleware"
	"TP-Back-Planity/web/models"
)

func (h *Handler) RequestAddEstablishment() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		token := request.Header.Get("Authorization")
		claims, err := middleware.DecodeJWT(token)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		professionalID := int(claims["user_id"].(float64))

		item := models.Request{}
		err = json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		item.ProfessionalID = professionalID

		data, err := h.Store.Request.RequestAddEstablishment(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status string         `json:"status"`
			Data   models.Request `json:"data"`
		}{
			Status: "success",
			Data:   data,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
