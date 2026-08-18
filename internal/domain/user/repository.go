package user

import "context"

type Repository interface {
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	AddUser(ctx context.Context, inputUser *User) (int, error)
	GetUserByID(ctx context.Context, userID int) (*User, error)
}
