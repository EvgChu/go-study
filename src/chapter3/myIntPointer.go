package main

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat = 98.5 // in c this dont work
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func negate(myBoolean *bool){
	*myBoolean = !(*myBoolean)
	return 
}

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println(*myIntPointer)
	fmt.Println(*createPointer())
	boolVar := true
	printPointer(&boolVar)
	negate(&boolVar)
	printPointer(&boolVar)
}
