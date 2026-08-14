package routers

import (
	"github.com/go-chi/chi/v5"
)

type Option func(router chi.Router)


func GetRouter(options ...Option) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		for _, option := range options {
			option(r)
		}
	})

	return router
}