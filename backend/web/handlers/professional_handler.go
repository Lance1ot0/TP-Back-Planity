package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"TP-Back-Planity/web/models"
	"TP-Back-Planity/web/utils"

	"github.com/go-chi/chi"
)

func (h *Handler) GetProfessional() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Je vais sortir un JSON, je rajoute le header correspondant
		writer.Header().Set("Content-Type", "application/json")

		professionals, _ := h.Store.Professional.GetProfessional()
		err := json.NewEncoder(writer).Encode(professionals)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetProfessionalById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		professional, _ := h.Store.Professional.GetProfessionalById(id)
		err := json.NewEncoder(writer).Encode(professional)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetProfessionalByEmail() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryEmail := chi.URLParam(request, "email")

		professional, _ := h.Store.Professional.GetProfessionalByEmail(QueryEmail)
		err := json.NewEncoder(writer).Encode(professional)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) AddProfessional() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		item := models.Professional{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		item.Password, err = utils.HashString(item.Password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := h.Store.Professional.AddProfessional(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status         string `json:"status"`
			ProfessionalID int    `json:"professionalID"`
		}{
			Status:         "succes",
			ProfessionalID: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
