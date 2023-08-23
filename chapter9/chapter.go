package main

import (
	"fmt"
	"hfg/chapter9/calendar"
	"log"
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

func (n *Number) Display() {
	fmt.Println(*n)
}

func taskPage315() {
	number := Number(4)
	number.Double()
	number.Display()
}

func examplePage338() {
	date := calendar.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())
	fmt.Println(date.Month())
	fmt.Println(date.Day())
}

func examplePage347() {
	event := calendar.Event{}
	err := event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(5)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Month())
	fmt.Println(event.Day())
}

func main() {
	examplePage309()
	taskPage312()
	taskPage315()
	examplePage338()
	examplePage347()
}
