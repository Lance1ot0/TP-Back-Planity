package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h *Handler) GetClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Je vais sortir un JSON, je rajoute le header correspondant
		writer.Header().Set("Content-Type", "application/json")

		todos, _ := h.Store.Client.GetClient()
		err := json.NewEncoder(writer).Encode(todos)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetClientById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		todos, _ := h.Store.Client.GetClientById(id)
		err := json.NewEncoder(writer).Encode(todos)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
