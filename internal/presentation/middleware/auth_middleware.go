package middleware

import (
	"net/http"
	"project/internal/presentation/appcontext"
	"project/internal/presentation/response"
	"project/internal/presentation/schemas"
	"strings"
)

type JWTManagerInterface interface {
	NewAccessToken(userID int) (string, error)
	NewRefreshToken(userID int) (string, error)
	ParseToken(inputToken string) (int, error)
}

func AuthMiddleware(jwtManager JWTManagerInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			token := req.Header.Get("Authorization")
			if !strings.HasPrefix(token, "Bearer ") {
				errorResponse := schemas.ErrorSchema{Error: "Unauthorized"}
				response.Error(w, http.StatusUnauthorized, errorResponse)
				return
			}

			userID, err := jwtManager.ParseToken(strings.TrimPrefix(token, "Bearer "))
			if err != nil {
				errorResponse := schemas.ErrorSchema{Error: "Unauthorized"}
				response.Error(w, http.StatusUnauthorized, errorResponse)
				return
			}
			ctx := appcontext.ContextWithUserID(req.Context(), userID)
			req = req.WithContext(ctx)

			next.ServeHTTP(w, req)
		})
	}
}
