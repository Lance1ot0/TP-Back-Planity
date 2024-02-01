package handlers

import (
	"encoding/json"
	"net/http"
	"TP-Back-Planity/web/models"
	"TP-Back-Planity/web/utils"
	"github.com/go-chi/chi"
	
)

func (h *Handler) GetAdminByEmail() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryEmail := chi.URLParam(request, "email")

		administrator, _ := h.Store.Admin.GetAdminByEmail(QueryEmail)
		err := json.NewEncoder(writer).Encode(administrator)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) LoginAdmin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		
		writer.Header().Set("Content-Type", "application/json")

		item := models.Administrator{}
		err := json.NewDecoder(request.Body).Decode(&item)

		email := item.Email
		password := item.Password

		item.Password, err = utils.HashString(password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		success, err := h.Store.Admin.LoginAdmin(email, password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		response := map[string]interface{}{"success": success}
		err = json.NewEncoder(writer).Encode(response)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
