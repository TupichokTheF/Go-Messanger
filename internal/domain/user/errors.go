package user

type BaseError struct {
	err string
}

func (be *BaseError) Error() string {
	return be.err
}

type IncorrectValue struct {
	value string
	BaseError
}
