package solutions

import (
	"fmt"
)

// Task01 - Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре
// Action от родительской структуры Human (аналог наследования).
func Task01() {
	action := Action{}
	action.sayHello()
}

// Human - parent struct
type Human struct {
	Name string
	Age  int
}

// sayHello - print hello
func (c *Human) sayHello() {
	fmt.Println("hello from human")
}

// Action - child struct from Human
type Action struct {
	Human
	ActionName string
}
