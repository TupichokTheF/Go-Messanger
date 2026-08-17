package schemas

type UserCreatedSchema struct {
	UserID int `json:"user_id" example:"1"`
}

type CreateUserSchema struct {
	Username string `json:"username" example:"maximEZ"`
	Email string `json:"email" example:"maxim@mail.ru"` 
	Password string `json:"password" example:"1Q2w3e"`
}

type AuthorizeSchema struct {
	Username string `json:"username" example:"maximEZ"`
	Password string `json:"password" example:"1Q2w3e"`
}

type TokensSchema struct {
	AccessToken string `json:"access_token"`
}