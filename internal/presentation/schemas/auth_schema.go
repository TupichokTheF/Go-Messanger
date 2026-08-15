package schemas

type UserCreatedSchema struct {
	UserID int `json:"user_id" example:"1"`
}

type CreateUserSchema struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}