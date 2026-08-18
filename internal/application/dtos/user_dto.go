package dtos

type UserCreateDTO struct {
	UserName     string
	UserPassword string
	UserEmail    string
}

type UserCreatedDTO struct {
	UserId int
}

type AuthorizeDTO struct {
	Username string
	Password string
}

type TokensDTO struct {
	AccessToken  string
	RefreshToken string
}

type UserInfoDTO struct {
	UserID    int
	Username  string
	UserEmail string
}
