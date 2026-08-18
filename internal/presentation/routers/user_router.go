package routers

import (
	"project/internal/presentation/handlers"
	"project/internal/presentation/middleware"

	"github.com/go-chi/chi/v5"
)

type JWTManagerInterface interface {
	NewAccessToken(userID int) (string, error)
	NewRefreshToken(userID int) (string, error)
	ParseToken(inputToken string) (int, error)
}

func WithUserRouter(handler *handlers.UserHandler, jwtManager JWTManagerInterface) Option {
	return func(router chi.Router) {
		router.Route("/user", func(r chi.Router) {
			r.Use(middleware.AuthMiddleware(jwtManager))

			r.Get("/me", handler.UserInfo)
		})
	}
}
