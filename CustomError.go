package zcb

type CustomError struct {
	msg string
}

func(customError CustomError) Error() string  {
	return customError.msg
}
