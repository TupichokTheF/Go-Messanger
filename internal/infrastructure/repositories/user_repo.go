package repositories

import (
	"context"
	"project/internal/domain/user"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	*pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		Pool: pool,
	}
}

func (userRepo *UserRepository) GetUserById(userId int) (*user.User, error) {
	var u *user.UserState = new(user.UserState)
	err := userRepo.Pool.QueryRow(context.TODO(),
		"SELECT * FROM users WHERE user_id = $1",
		userId).Scan(&u.ID, &u.UserName, &u.UserEmail, &u.UserPassword)

	if err != nil {
		return nil, user.NotFoundError
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
		return 0, user.AlreadyExistError
	}

	return id, nil
}
