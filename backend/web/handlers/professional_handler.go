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

func (h *Handler) GetHairSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		token := request.Header.Get("Authorization")
		claims, err := middleware.DecodeJWT(token)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id := int(claims["user_id"].(float64))

		hairSalon, _ := h.Store.Professional.GetHairSalon(id)
		if hairSalon.HairSalonID == 0 {
			http.Error(writer, "No hair salon found", http.StatusNotFound)
			return
		}

		err = json.NewEncoder(writer).Encode(hairSalon)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) GetRequest() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		token := request.Header.Get("Authorization")
		claims, err := middleware.DecodeJWT(token)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id := int(claims["user_id"].(float64))

		requests, _ := h.Store.Professional.GetRequest(id)
		if requests.RequestID == 0 {
			http.Error(writer, "No request found", http.StatusNotFound)
			return
		}

		err = json.NewEncoder(writer).Encode(requests)
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

		_, err = h.Store.Professional.AddProfessional(item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status bool `json:"status"`
		}{
			Status: true,
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
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Authentication failed"))
			return
		}

		password, err := h.Store.Professional.GetPasswordHash(professional.ProfessionalID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = utils.CompareHashAndPassword(password, item.Password)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Authentication failed"))
			return
		}

		token, err := middleware.GenerateJWT(professional.ProfessionalID, "professional")
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
				Role:   "professional",
				Token:  token,
			})
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func (h *Handler) GetEmployeeAvailability() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		employee, _ := h.Store.Professional.GetEmployeeAvailability(id)
		err := json.NewEncoder(writer).Encode(employee)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) AddEmployeeAvailability() http.HandlerFunc {
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

		_, err = h.Store.Professional.AddEmployeeAvailability(id, item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status bool `json:"status"`
		}{
			Status: true,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
