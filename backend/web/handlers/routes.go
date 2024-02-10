package handlers

import (
	database "TP-Back-Planity/web/store"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandler(store *database.Store) *chi.Mux {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("TP-BACK-PLANITY API"))
	})

	handler.Route("/api", func(r chi.Router) {
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("API"))
		})

		r.Route("/client", func(r chi.Router) {
			r.Get("/", handler.GetClient())
			r.Get("/{id}", handler.GetClientById())
			r.Post("/hairSalon", handler.GetHairSalon())
			r.Post("/register", handler.AddClient())
			r.Post("/login", handler.LoginClient())
			r.Post("/reservation", handler.AddReservation())
			r.Get("/{clientId}/reservations", handler.ListReservations())
		})

		r.Route("/admin", func(r chi.Router) {
			// r.Get("/", handler.GetAdmin())
			// r.Get("/{id}", handler.GetAdminById())
			r.Get("/email/{email}", handler.GetAdminByEmail())
			r.Post("/login", handler.LoginAdmin())
			r.Get("/requests", handler.ListRequests())
			r.Put("/request/{id}", handler.HandleRequest())
		})

		r.Route("/professional", func(r chi.Router) {
			r.Get("/", handler.GetProfessional())
			r.Get("/{id}", handler.GetProfessionalById())
			r.Get("/email/{email}", handler.GetProfessionalByEmail())
			r.Post("/register", handler.AddProfessional())
			r.Post("/login", handler.LoginProfessional())
			r.Post("/employee", handler.AddEmploye())
			r.Post("/service", handler.AddService())
			r.Post("/request", handler.RequestAddEstablishment())
		})
	})

	return handler.Mux
}

type Handler struct {
	*chi.Mux
	*database.Store
}
