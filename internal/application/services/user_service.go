package services

import (
	"context"
	"fmt"
	"project/internal/application/dtos"
	"project/internal/domain/user"
)


type UserService struct {
	userRepo user.Repository
}

func NewUserService(userRepo user.Repository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (userService *UserService) CreateNewUser(ctx context.Context, userCreateDTO dtos.UserCreateDTO) (*dtos.UserCreatedDTO, error) {
	user, err := user.New(userCreateDTO.UserName, userCreateDTO.UserEmail, userCreateDTO.UserPassword)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	
	userId, err := userService.userRepo.AddUser(user)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return new(dtos.UserCreatedDTO{
		UserId: userId,
	}), nil
}