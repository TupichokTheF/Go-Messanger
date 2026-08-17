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

func NewAuthHandler(userService ports.UserService) *AuthHander {
	return &AuthHander{
		userService: userService,
	}
}

// @Summary  Регистрация пользователя
// @Tags     auth
// @Accept   json
// @Produce  json
// @Param    request body     schemas.CreateUserSchema true "Данные пользователя"
// @Success  201     {object} schemas.UserCreatedSchema
// @Failure  400     {object} schemas.ErrorSchema
// @Router   /auth/register [post]
func (handler *AuthHander) CreateUser(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var request schemas.CreateUserSchema
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		errorResponse := schemas.ErrorSchema{Error: "Invalid request body"} 
		response.Error(w, http.StatusBadRequest, errorResponse)
		return
	}

	result, err := handler.userService.CreateNewUser(req.Context(), mappers.FromCreatedSchemaToDTO(&request))
	if err != nil {
		status, errorMessage := mappers.FromApplicationToApiError(err)
		errorResponse := schemas.ErrorSchema{Error: errorMessage}
		response.Error(w, status, errorResponse)
		return
	}

	response.JSON(w, http.StatusCreated, mappers.FromCreatedDTOToSchema(result))
}

// @Summary  Авторизация пользователя
// @Tags     auth
// @Accept   json
// @Produce  json
// @Param    request body     schemas.AuthorizeSchema true "Учётные данные"
// @Success  200     {object} schemas.TokensSchema
// @Failure  400     {object} schemas.ErrorSchema
// @Failure  401     {object} schemas.ErrorSchema
// @Router   /auth/login [post]
func (handler *AuthHander) Authorization(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var request schemas.AuthorizeSchema
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		errorResponse := schemas.ErrorSchema{Error: "Invalid request body"} 
		response.Error(w, http.StatusBadRequest, errorResponse)
		return
	}

	result, err := handler.userService.AuthorizeUser(req.Context(), mappers.FromAuthorizeSchemaToDTO(&request))
	if err != nil {
		status, errorMessage := mappers.FromApplicationToApiError(err)
		errorResponse := schemas.ErrorSchema{Error: errorMessage}
		response.Error(w, status, errorResponse)
		return
	}
	responseOptons := []response.ResponseOption{
		response.WithRefreshTokenCookie(result.RefreshToken),
	}
	response.JSON(w, http.StatusCreated, mappers.FromTokensDTOToSchema(result), responseOptons...)
}