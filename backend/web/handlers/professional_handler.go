package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"TP-Back-Planity/web/middleware"
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
			Status:         "success",
			ProfessionalID: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) LoginProfessional() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		item := models.Professional{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		professional, err := h.Store.Professional.GetProfessionalByEmail(item.Email)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if professional.ProfessionalID == 0 {
			err = json.NewEncoder(writer).Encode(struct {
				Status string `json:"status"`
				Error  string `json:"error"`
			}{
				Status: "error",
				Error:  "Email not found",
			})
			return
		}

		password, err := h.Store.Professional.GetPasswordHash(professional.ProfessionalID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = utils.CompareHashAndPassword(password, item.Password)
		if err != nil {
			err = json.NewEncoder(writer).Encode(struct {
				Status string `json:"status"`
				Error  string `json:"error"`
			}{
				Status: "error",
				Error:  "Password incorrect",
			})
			return
		}

		token, err := middleware.GenerateJWT(professional.ProfessionalID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		} else {
			err = json.NewEncoder(writer).Encode(struct {
				Status string `json:"status"`
				Token  string `json:"token"`
			}{
				Status: "success",
				Token:  token,
			})
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
