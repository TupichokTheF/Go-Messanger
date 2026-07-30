package user

type User struct {
	Id           int
	UserName     UserName
	UserEmail    UserEmail
	UserPassword UserPassword
}

func CreateUser(userName_, userEmail_, userPassword_ string) (*User, error) {
	userName, err := CreateUserName(userName_)
	if err != nil {
		return nil, err
	}

	userPass, err := CreatePassword(userPassword_)
	if err != nil {
		return nil, err
	}

	userEmail, err := CreateUserEmail(userEmail_)
	if err != nil {
		return nil, err
	}

	return new(User{
		UserName:     userName,
		UserPassword: userPass,
		UserEmail:    userEmail,
	}), nil
}
