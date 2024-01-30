package handlers

import (
	database "TP-Back-Planity/web/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandler(store *database.Store) *chi.Mux {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Get("/", handler.GetClient())

	handler.Route("/api", func(r chi.Router) {
		r.Get("/client", handler.GetClient())
		r.Get("/client/{id}", handler.GetClientById())
	})

	return handler.Mux
}

type Handler struct {
	*chi.Mux
	*database.Store
}
