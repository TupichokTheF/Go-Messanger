package dtos


type UserCreateDTO struct {
	UserName string
	UserPassword string
	UserEmail string
}

type UserCreatedDTO struct {
	UserId int
}