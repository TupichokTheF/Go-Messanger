package user


type Repository interface {
	GetUserById(userId int) (*User, error)
	AddUser(inputUser *User) (int, error)
}