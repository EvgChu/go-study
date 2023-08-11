package main

import (
	"fmt"
	"hfg/chapter5/datafile"
	"log"
)

func main() {
	numbers, err := datafile.GetFloats("chapter5/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
