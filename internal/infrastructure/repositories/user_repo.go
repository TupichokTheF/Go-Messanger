package repositories

import (
	"context"
	"fmt"
	"project/internal/domain/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	*pgxpool.Pool
}

func CreateNewUserRepository(pool *pgxpool.Pool) (*UserRepository, error) {
	return &UserRepository{
		Pool: pool,
	}, nil
}

func (userRepo *UserRepository) GetUserById(userId int) (*user.User, error) {
	var u *user.UserState = new(user.UserState)
	err := userRepo.Pool.QueryRow(context.TODO(),
		"SELECT * FROM users WHERE user_id = $1",
		userId).Scan(&u.ID, &u.UserName, &u.UserEmail, &u.UserPassword)

	if err != nil {
		return nil, fmt.Errorf("Oshibochka NE LOL")
	}

	return user.Reconstitute(*u), nil
}

func (userRepo *UserRepository) AddUser(inputUser *user.User) (int, error) {
	var id int
	userState := inputUser.State()
	err := userRepo.Pool.QueryRow(context.TODO(),
		`INSERT INTO users (username, email, password) 
		VALUES ($1, $2, $3)
		RETURNING user_id`,
		userState.UserName, userState.UserEmail, userState.UserPassword).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("Error while adding user: %v", err)
	}

	return id, nil
}
