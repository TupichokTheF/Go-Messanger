package user


type UserRepoInterface interface {
	GetUserById(userId int) (*User, error)
	AddUser(inputUser *User) (int, error)
}