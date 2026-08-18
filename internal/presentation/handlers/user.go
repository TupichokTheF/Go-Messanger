package handlers

import (
	"net/http"
	"project/internal/application/ports"
	"project/internal/presentation/appcontext"
	"project/internal/presentation/mappers"
	"project/internal/presentation/response"
	"project/internal/presentation/schemas"
)

type UserHandler struct {
	userService ports.UserService
}

func NewUserHandlers(userService ports.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) UserInfo(w http.ResponseWriter, req *http.Request) {
	userID, ok := appcontext.UserIDFromContext(req.Context())
	if !ok {
		errorResponse := schemas.ErrorSchema{Error: "Unauthorized"}
		response.Error(w, http.StatusUnauthorized, errorResponse)
		return
	}

	result, err := handler.userService.GetUserInfo(req.Context(), userID)
	if err != nil {
		status, errorMessage := mappers.FromApplicationToApiError(err)
		errorResponse := schemas.ErrorSchema{Error: errorMessage}
		response.Error(w, status, errorResponse)
		return
	}

	response.JSON(w, http.StatusOK, mappers.FromUserInfoDTOToSchema(result))
}
