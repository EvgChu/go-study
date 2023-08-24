package main

import (
	"fmt"
)

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {

	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {

	return "Hi from MethodWithReturnValue"
}

func (my MyType) MethodNotInInterface() {

	fmt.Println("MethodNotInInterface called")
}

func examplePage359() {
	var value MyInterface
	value = MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}

func main() {
	examplePage359()
}
