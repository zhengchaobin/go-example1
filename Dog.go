package zcb

import "fmt"

type Dog struct {
	Name string
}

func (dog Dog) Sleep(){
	fmt.Println(dog.Name + "睡觉")
}

func (dog Dog) Move(){
	fmt.Println(dog.Name + "移动")
}

