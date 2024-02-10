package handlers

import (
	"TP-Back-Planity/web/middleware"
	"TP-Back-Planity/web/models"
	"TP-Back-Planity/web/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

		id, err := h.Store.Admin.LoginAdmin(email, password)
		if err != nil {

			if id == 0 {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte("Authentication failed"))
				// err = json.NewEncoder(writer).Encode(struct {
				// 	Status string `json:"status"`
				// 	Error  string `json:"error"`
				// }{
				// 	Status: "error",
				// 	Error:  "Authentication failed",
				// })
				return
			}

			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := middleware.GenerateJWT(id, "admin")
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

func (h *Handler) HandleRequest() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var status string

		item := models.Request{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		status = item.RequestStatus
		writer.Header().Set("Content-Type", "application/json")
		QueryId := chi.URLParam(request, "id")

		id, _ := strconv.Atoi(QueryId)

		requestSalon, err := h.Store.Admin.UpdateRequest(id, status)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		if !requestSalon {
			http.Error(writer, "Invalid request", http.StatusBadRequest)
			return
		}

		if status == "accepted" {

			acceptedRequest, err := h.Store.Admin.GetRequestById(id)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			result, _ := h.Store.Admin.CreateSalon(acceptedRequest)
			err = json.NewEncoder(writer).Encode(result)
		} else {
			err = json.NewEncoder(writer).Encode(status)
		}

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetHairSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var name string

		item := models.HairSalon{}
		err := json.NewDecoder(request.Body).Decode(&item)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		name = item.Name
		fmt.Println(name)
		writer.Header().Set("Content-Type", "application/json")

		hairSalon, err := h.Store.Client.GetHairSalon(name)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(writer).Encode(hairSalon)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
