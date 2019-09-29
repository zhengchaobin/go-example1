package zcb

import "testing"

func TestDog_Move(t *testing.T) {
	var animal Animal
	animal = Dog{"小狗"}
	animal.Sleep()
	animal.Move()
}