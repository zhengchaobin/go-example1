package zcb

import "fmt"

type Cat struct {
	// 要跨包调用的属性首字母必须大写
	// 可理解为大写即是共有的，小写即是私有的
	Name string
}

func (cat Cat) Sleep()  {
	fmt.Println(cat.Name + "睡觉")
}

func (cat Cat) Move() {
	fmt.Println(cat.Name + "移动")
}

type Bird struct {
	// 匿名字段，那么默认Bird就包含了Cat的所有字段
	Cat
	Name string
}