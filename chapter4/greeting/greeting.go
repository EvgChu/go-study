// Package greeting for test go doc
package greeting

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

func Hi() {
	fmt.Println("Hi!")
}

// Say shot and long hello
func AllGreetings() {
	Hello()
	Hi()
}
