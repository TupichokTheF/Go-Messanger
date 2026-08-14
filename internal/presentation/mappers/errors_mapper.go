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
		return http.StatusBadRequest, appError.Error()
	case errors.Is(appError, user.AlreadyExistError):
		return http.StatusConflict, appError.Error()
	case errors.Is(appError, user.NotFoundError):
		return http.StatusBadRequest, appError.Error()
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}