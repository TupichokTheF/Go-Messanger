package user

type User struct {
	id           int
	userName     UserName
	userEmail    UserEmail
	userPassword UserPassword
}

type HasherInterface interface {
	Hash(password string) (string, error)
	Verify(password, hash string) bool
}

func New(inputUserName, inputUserEmail, inputUserPassword string, hasher HasherInterface) (*User, error) {
	userName, err := CreateUserName(inputUserName)
	if err != nil {
		return nil, err
	}

	_, err = CreatePassword(inputUserPassword)
	if err != nil {
		return nil, err
	}
	hashedPass, err := hasher.Hash(inputUserPassword)
	if err != nil {
		return nil, InvalidPassword
	}

	userEmail, err := CreateUserEmail(inputUserEmail)
	if err != nil {
		return nil, err
	}

	return &User{
		userName:     userName,
		userPassword: UserPassword{value: hashedPass},
		userEmail:    userEmail,
	}, nil
}

func (u *User) Username() UserName {
	return u.userName
}

func (u *User) Email() UserEmail {
	return u.userEmail
}

func (u *User) Password() UserPassword {
	return u.userPassword
}

func (u *User) ID() int {
	return u.id
}

func (u *User) VerifyPassword(raw string, hasher HasherInterface) bool {
	return hasher.Verify(raw, u.userPassword.value)
}
