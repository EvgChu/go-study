package main

import (
	"fmt"
	"log"
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
	horn, ok := noiseMaker.(Horn)
	if ok {
		horn.MakeSound()
	}
}

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func examplePage375() {
	var err error = checkTemperature(21.379, 100.0)
	if err != nil {
		log.Fatal(err)
	}
}

type CoffeePot string

func (c CoffeePot) String() string {
	return string(c) + " coffee pot"
}

func examplePage376() {
	coffeePot := CoffeePot("LuxBrew")
	fmt.Println(coffeePot.String())
	fmt.Println(coffeePot)
	fmt.Printf("%s", coffeePot)
}

func AcceptAnything(thing interface{}) {
}

func examplePage379() {
	AcceptAnything(3.1415)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("Toyco Canary"))
}

func examplePage387() {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	fmt.Println("Nice weather, eh?")
}

func main() {
	examplePage359()
	examplePage362()
	examplePage369()
	examplePage375()
	examplePage376()
	examplePage379()
	examplePage387()
}
