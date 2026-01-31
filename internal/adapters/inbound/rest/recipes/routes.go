package recipes

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, handler *Handler) {
	r.Route("/recipes", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetByID)
		r.Post("/", handler.Create)
		r.Delete("/{id}", handler.Delete)
	})
}
