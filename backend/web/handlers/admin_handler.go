package handlers

import (
	"encoding/json"
	"TP-Back-Planity/web/models"
	"TP-Back-Planity/web/utils"
	"TP-Back-Planity/web/middleware"
	"github.com/go-chi/chi"
	"net/http"
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

		id, err := h.Store.Admin.LoginAdmin(email, password)
		if err != nil{

			if id == 0 {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte("Authentication failed"))
				return 
			}

			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := middleware.GenerateJWT(id)
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

func (h *Handler) ListRequests() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		writer.Header().Set("Content-Type", "application/json")

		requests, _ := h.Store.Admin.ListRequests()
		err := json.NewEncoder(writer).Encode(requests)
		
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
