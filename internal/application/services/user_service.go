package services

import (
	"context"
	"fmt"
	"project/internal/application/dtos"
	"project/internal/domain/user"
)

type UserService struct {
	userRepo   user.Repository
	jwtManager JWTManagerInterface
	hasher     HasherInterface
}

type JWTManagerInterface interface {
	NewAccessToken(userID int) (string, error)
	NewRefreshToken(userID int) (string, error)
	ParseToken(inputToken string) (int, error)
}

type HasherInterface interface {
	Hash(password string) (string, error)
	Verify(password, hash string) bool
}

func NewUserService(userRepo user.Repository, jwtManager JWTManagerInterface, hasher HasherInterface) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
		hasher:     hasher,
	}
}

func (userService *UserService) CreateNewUser(ctx context.Context, userCreateDTO *dtos.UserCreateDTO) (*dtos.UserCreatedDTO, error) {
	createdUser, err := user.New(userCreateDTO.UserName, userCreateDTO.UserEmail, userCreateDTO.UserPassword, userService.hasher)
	if err != nil {
		return nil, fmt.Errorf("Create user: %w", err)
	}

	userID, err := userService.userRepo.AddUser(ctx, createdUser)
	if err != nil {
		return nil, fmt.Errorf("Create user: %w", err)
	}

	return &dtos.UserCreatedDTO{
		UserId: userID,
	}, nil
}

func (userService *UserService) AuthorizeUser(ctx context.Context, authorizeDTO *dtos.AuthorizeDTO) (*dtos.TokensDTO, error) {
	u, err := userService.userRepo.GetUserByUsername(ctx, authorizeDTO.Username)
	if err != nil {
		return nil, fmt.Errorf("User authorization: %w", err)
	}

	if ok := u.VerifyPassword(authorizeDTO.Password, userService.hasher); !ok {
		return nil, fmt.Errorf("User Authorization: %w", user.InvalidPassword)
	}

	accessToken, err := userService.jwtManager.NewAccessToken(u.ID())
	if err != nil {
		return nil, fmt.Errorf("User authorization: %w", err)
	}

	refreshToken, err := userService.jwtManager.NewRefreshToken(u.ID())
	if err != nil {
		return nil, fmt.Errorf("User authorization: %w", err)
	}

	return &dtos.TokensDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
