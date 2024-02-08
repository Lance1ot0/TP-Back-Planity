package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"TP-Back-Planity/web/models"

	"github.com/go-chi/chi"
)

func (h *Handler) AddAvailability() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		item := models.Availability{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		availability, err := h.Store.Availability.AddAvailability(id, item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status              string `json:"status"`
			models.Availability `json:"data"`
		}{
			Status:       "success",
			Availability: availability,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
