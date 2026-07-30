package user

import (
	"net/mail"
	"strings"
	"unicode"
)

type (
	UserName     string
	UserPassword string
	UserEmail    string
)

func CreateUserName(value string) (UserName, error) {
	if len(value) < 4 {
		return "", &IncorrectValue{BaseError: BaseError{err: "Username is too short"}, value: value}
	} else if len(value) > 30 {
		return "", &IncorrectValue{BaseError: BaseError{err: "Username is too long"}, value: value}
	} else if strings.TrimSpace(value) == "" {
		return "", &IncorrectValue{BaseError: BaseError{err: "Empty username"}, value: value}
	}

	return UserName(value), nil
}

func CreatePassword(value string) (UserPassword, error) {
	if len(value) < 6 {
		return "", &IncorrectValue{BaseError: BaseError{err: "Password is too short"}}
	} else if len(value) > 30 {
		return "", &IncorrectValue{BaseError: BaseError{err: "Password is too long"}}
	} else if strings.ToLower(value) == value {
		return "", &IncorrectValue{BaseError: BaseError{err: "Password required upper symbol"}}
	} else if strings.ToUpper(value) == value {
		return "", &IncorrectValue{BaseError: BaseError{err: "Password required lower symbol"}}
	}
	hasDigit := func() bool {
		for _, r := range value {
			if unicode.IsDigit(r) {
				return true
			}
		}
		return false
	}
	if !hasDigit() {
		return "", &IncorrectValue{BaseError: BaseError{err: "Password required digit"}}
	}

	return UserPassword(value), nil
}

func CreateUserEmail(value string) (UserEmail, error) {
	if len(value) > 200 {
		return "", &IncorrectValue{BaseError: BaseError{err: "Email address is too long"}, value: value}
	}

	_, err := mail.ParseAddress(value)
	if err != nil {
		return "", &IncorrectValue{BaseError: BaseError{err: "Incorrect email address"}, value: value}
	}

	return UserEmail(value), nil
}
