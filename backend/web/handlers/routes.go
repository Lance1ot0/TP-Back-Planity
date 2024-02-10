package handlers

import (
	mid "TP-Back-Planity/web/middleware"
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
			r.With(mid.JWTMiddleware).Get("/", handler.GetClient())
			r.With(mid.JWTMiddleware).Get("/{id}", handler.GetClientById())

			r.Post("/register", handler.AddClient())
			r.Post("/login", handler.LoginClient())
			r.With(mid.JWTMiddleware).Post("/hairSalon", handler.GetHairSalon())
			r.With(mid.JWTMiddleware).Post("/reservation", handler.AddReservation())
		})

		r.Route("/admin", func(r chi.Router) {
			r.With(mid.JWTMiddleware).Get("/email/{email}", handler.GetAdminByEmail())
			r.With(mid.JWTMiddleware).Get("/requests", handler.ListRequests())

			r.Post("/login", handler.LoginAdmin())

			r.Put("/request/{id}", handler.HandleRequest())
		})

		r.Route("/professional", func(r chi.Router) {
			r.With(mid.JWTMiddleware).Get("/", handler.GetProfessional())
			r.With(mid.JWTMiddleware).Get("/{id}", handler.GetProfessionalById())
			r.With(mid.JWTMiddleware).Get("/email/{email}", handler.GetProfessionalByEmail())

			r.Post("/register", handler.AddProfessional())
			r.Post("/login", handler.LoginProfessional())
			r.With(mid.JWTMiddleware).Post("/employee", handler.AddEmploye())
			r.With(mid.JWTMiddleware).Post("/service", handler.AddService())
			r.With(mid.JWTMiddleware).Post("/request", handler.RequestAddEstablishment())
		})
	})

	return handler.Mux
}

type Handler struct {
	*chi.Mux
	*database.Store
}
