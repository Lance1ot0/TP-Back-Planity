package handlers

import (
	mdw "TP-Back-Planity/web/middleware"
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
			r.With(mdw.Authorization).Get("/", handler.GetClient())
			r.With(mdw.Authorization).Get("/{id}", handler.GetClientById())

			r.Post("/register", handler.AddClient())
			r.Post("/login", handler.LoginClient())
			r.With(mdw.Authorization).Post("/hairSalon", handler.ResearchHairSalon())
			r.With(mdw.Authorization).Post("/reservation", handler.AddReservation())
			r.Post("/reservation", handler.AddReservation())
			r.Get("/{clientId}/reservations", handler.ListReservations())
			r.Put("/reservation/{reservationId}", handler.CancelReservation())
		})

		r.Route("/admin", func(r chi.Router) {
			r.With(mdw.Authorization).Get("/email/{email}", handler.GetAdminByEmail())
			r.With(mdw.Authorization).Get("/requests", handler.ListRequests())

			r.Post("/login", handler.LoginAdmin())

			r.Put("/request/{id}", handler.HandleRequest())
		})

		r.Route("/professional", func(r chi.Router) {
			r.With(mdw.Authorization).Get("/", handler.GetProfessional())
			r.With(mdw.Authorization).Get("/{id}", handler.GetProfessionalById())
			r.With(mdw.Authorization).Get("/email/{email}", handler.GetProfessionalByEmail())
			r.With(mdw.Authorization).Get("/hairSalon", handler.GetHairSalon())
			r.With(mdw.Authorization).Get("/request", handler.GetRequest())
			r.With(mdw.Authorization).Get("/employee/{id}", handler.GetAllEmployee())
			r.With(mdw.Authorization).Get("/employee/availability/{id}", handler.GetEmployeeAvailability())

			r.Post("/register", handler.AddProfessional())
			r.Post("/login", handler.LoginProfessional())
			r.With(mdw.Authorization).Post("/employee", handler.AddEmploye())
			r.With(mdw.Authorization).Post("/service", handler.AddService())
			r.With(mdw.Authorization).Post("/request", handler.RequestAddEstablishment())
			r.With(mdw.Authorization).Post("/employee/availability/{id}", handler.AddEmployeeAvailability())
		})
	})

	return handler.Mux
}

type Handler struct {
	*chi.Mux
	*database.Store
}
