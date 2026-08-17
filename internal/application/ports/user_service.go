package ports

import (
	"context"
	"project/internal/application/dtos"
)


type UserService interface {
	CreateNewUser(ctx context.Context, userCreateDTO *dtos.UserCreateDTO) (*dtos.UserCreatedDTO, error)
	AuthorizeUser(ctx context.Context, authorizeDTO *dtos.AuthorizeDTO) (*dtos.TokensDTO, error)
}