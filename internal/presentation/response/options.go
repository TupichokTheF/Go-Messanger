package response

import "net/http"


type ResponseOption func(w http.ResponseWriter)


func WithRefreshTokenCookie(token string) ResponseOption {
	return func(w http.ResponseWriter) {
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    token,
			HttpOnly: true,                
			Secure:   true,
		})
	}
}