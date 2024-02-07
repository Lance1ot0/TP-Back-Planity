package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handler) GetReservation() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		reservations, _ := h.Store.Reservation.GetAllReservations(id)
		err := json.NewEncoder(writer).Encode(reservations)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}