package routers

import (
	"project/internal/presentation/handlers"

	"github.com/go-chi/chi/v5"
)


func WithAuthRouter(handler handlers.AuthHander) Option {
	return func(router chi.Router) {
		router.Route("/auth", func(r chi.Router) {
			r.Post("/register", handler.CreateUser)
		})
	}
}
