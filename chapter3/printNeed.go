package main

import "fmt"

func printNeeded(width float64, height float64){
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main(){
	printNeeded(4.2, 3.0)
	printNeeded(4.0, 3.3)
	printNeeded(4, 3.6)
}