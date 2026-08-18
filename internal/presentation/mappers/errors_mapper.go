package mappers

import (
	"errors"
	"net/http"
	"project/internal/domain/user"
)

func FromApplicationToApiError(appError error) (int, string) {
	var validationError *user.ValidationError

	switch {
	case errors.As(appError, &validationError):
		return http.StatusBadRequest, validationError.Reason
	case errors.Is(appError, user.AlreadyExistError):
		return http.StatusConflict, "User already exists"
	case errors.Is(appError, user.NotFoundError):
		return http.StatusBadRequest, "User wasn't found"
	case errors.Is(appError, user.InvalidPassword):
		return http.StatusUnauthorized, "Invalid password"
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}
