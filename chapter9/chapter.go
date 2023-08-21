package main

import (
	"fmt"
)

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi", m)
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func examplePage309() {
	value := MyType("a MyType value")
	value.sayHi()
	value.MethodWithParameters(4, true)
	anotherValue := MyType("another value")
	anotherValue.sayHi()
}

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func (n *Number) Double() {
	*n *= 2
}

func taskPage312() {
	ten := Number(10)
	number := Number(2)
	ten.Add(4)
	ten.Subtract(5)
	number.Double()
	fmt.Println("number after calling Double:", number)
	four := Number(4)
	four.Add(3)
	four.Subtract(2)
}

func main() {
	examplePage309()
	taskPage312()
}
