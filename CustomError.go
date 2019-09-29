package zcb

type CustomError struct {
	msg string
}

func(customError CustomError) Error() string  {
	return customError.msg
}

func New(msgText string) *CustomError {
	return &CustomError{msg: msgText}
}
