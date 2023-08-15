package main

import (
	"fmt"
	"hfg/chapter5/datafile"
	"log"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func main() {
	numbers, err := datafile.GetFloats("chapter5/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Average: %0.2f\n", average(numbers...))
}
