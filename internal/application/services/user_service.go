package services

import (
	"context"
	"project/internal/application/dtos"
	errors "project/internal/application/errors"
	"project/internal/domain/user"
)


type UserService struct {
	userRepo user.UserRepoInterface
}

func CreateNewUserService(userRepo user.UserRepoInterface) (*UserService, error) {
	return &UserService{
		userRepo: userRepo,
	}, nil
}

func (userService *UserService) CreateNewUser(ctx context.Context, userCreateDTO dtos.UserCreateDTO) (*dtos.UserCreatedDTO, error) {
	user, err := user.CreateUser(userCreateDTO.UserName, userCreateDTO.UserEmail, userCreateDTO.UserPassword)
	if err != nil {
		return nil, &errors.IncorrectValue{BaseApplicationError: errors.BaseApplicationError{Err: err.Error()}}
	}
	
	userId, err := userService.userRepo.AddUser(user)
	if err != nil {
		return nil, &errors.UserAlreadyExist{BaseApplicationError: errors.BaseApplicationError{Err: err.Error()}}
	}

	return new(dtos.UserCreatedDTO{
		UserId: userId,
	}), nil
}