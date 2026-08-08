package applicationErrors


type BaseApplicationError struct {
	Err string
}

func (baseErr *BaseApplicationError) Error() string {
	return baseErr.Err
}

type IncorrectValue struct {
	value string
	BaseApplicationError
}

type UserAlreadyExist struct {
	BaseApplicationError
}