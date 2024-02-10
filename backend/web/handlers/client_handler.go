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

func (h *Handler) AddClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		item := models.Client{}
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

		id, err := h.Store.Client.AddClient(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status   string `json:"status"`
			ClientID int    `json:"clientID"`
		}{
			Status:   "success",
			ClientID: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) LoginClient() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		item := models.Client{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		client, err := h.Store.Client.GetClientByEmail(item.Email)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		if client.ClientID == 0 {
			err = json.NewEncoder(writer).Encode(struct {
				Status string `json:"status"`
				Error  string `json:"error"`
			}{
				Status: "error",
				Error:  "Email not found",
			})
			return
		}

		password, err := h.Store.Client.GetPasswordHash(client.ClientID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(password, item.Password)

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

		token, err := middleware.GenerateJWT(client.ClientID, "client")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		} else {
			err = json.NewEncoder(writer).Encode(struct {
				Status string `json:"status"`
				Role   string `json:"role"`
				Token  string `json:"token"`
			}{
				Status: "success",
				Role:   "client",
				Token:  token,
			})
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func (h *Handler) AddReservation() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		item := models.Reservation{}
		err := json.NewDecoder(request.Body).Decode(&item)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		reservation, err := h.Store.Client.AddReservation(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status      string             `json:"status"`
			Reservation models.Reservation `json:"data"`
		}{
			Status:      "success",
			Reservation: reservation,
		})
	}
}
