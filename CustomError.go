package zcb

type CustomError struct {
	msg string
}

func(customError CustomError) Error() string  {
	return customError.msg
}

//函数定义返回类型为指针类型时， if e != nil 会为true, 此时调用Error函数时会报空指针异常
//详见example2.ErrorTest.customErrorTest2
//func(customError *CustomError) Error() string  {
//	if  customError == nil {
//		return ""
//	}
//	return customError.msg
//}


func New(msgText string) CustomError {
	return CustomError{msg: msgText}
}

func New2(msgText string) *CustomError {
	return &CustomError{msg: msgText}
}
