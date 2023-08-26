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

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

func examplePage362() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}
func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

func examplePage369() {
	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}

func main() {
	examplePage359()
	examplePage362()
	examplePage369()
}
