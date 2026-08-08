package user

type User struct {
	Id           int
	UserName     UserName
	UserEmail    UserEmail
	UserPassword UserPassword
}

func CreateUser(inputUserName, inputUserEmail, inputUserPassword string) (*User, error) {
	userName, err := CreateUserName(inputUserName)
	if err != nil {
		return nil, err
	}

	userPass, err := CreatePassword(inputUserPassword)
	if err != nil {
		return nil, err
	}

	userEmail, err := CreateUserEmail(inputUserEmail)
	if err != nil {
		return nil, err
	}

	return &User{
		UserName:     userName,
		UserPassword: userPass,
		UserEmail:    userEmail,
	}, nil
}
