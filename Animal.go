package zcb

type Animal interface {
	// 要跨包调用的方法首字母必须大写
	Sleep()
	Move()
}