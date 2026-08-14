package handlers

import (
	"encoding/json"
	"net/http"
	"project/internal/application/ports"
	"project/internal/presentation/mappers"
	"project/internal/presentation/response"
	"project/internal/presentation/schemas"
)


type AuthHander struct {
	userService ports.UserService
}

func NewAuthHandler(userService ports.UserService) (*AuthHander, error) {
	return &AuthHander{
		userService: userService,
	}, nil
}


func (handler *AuthHander) CreateUser(w http.ResponseWriter, req *http.Request) {
	var request schemas.CreateUserSchema
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer req.Body.Close()

	result, err := handler.userService.CreateNewUser(req.Context(), *mappers.FromCreatedSchemaToDTO(&request))
	if err != nil {
		status, errorMessage := mappers.FromApplicationToApiError(err)
		response.Error(w, status, errorMessage)
		return
	}

	response.JSON(w, http.StatusCreated, result)
}