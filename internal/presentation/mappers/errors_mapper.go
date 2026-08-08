package mappers

import (
	"net/http"
	applicationErrors "project/internal/application/errors"
)


func FromApplicationToApiError(appError error) (int, string) {
	switch appError.(type) {
	case *applicationErrors.IncorrectValue:
		return http.StatusBadRequest, appError.Error()
	case *applicationErrors.UserAlreadyExist:
		return http.StatusConflict, appError.Error()
	default:
		return http.StatusInternalServerError, "Interal server error"
	}
}