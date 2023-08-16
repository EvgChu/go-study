package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	max := math.Inf(-1)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func sum(numbers ...int) int {
	var sum int = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.7, 98.5, 92.3))
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
	fmt.Println(sum(1, 3, 4, 5))
	jewelry := make(map[string]float64)
	jewelry["necklace"] = 89.99
	jewelry["earrings"] = 79.99
	clothing := map[string]float64{"pants": 59.99, "shirt": 39.99}
	fmt.Println("Earrings:", jewelry["earrings"])
	fmt.Println("Necklace:", jewelry["necklace"])
	fmt.Println("Shirt:", clothing["shirt"])
	fmt.Println("Pants:", clothing["pants"])
	counters := map[string]int{"a": 3, "b": 0}
	var value int
	var ok bool
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	value = counters["c"]
	fmt.Println(value, "c")
	value, ok = counters["c"]
	fmt.Println(value, ok)
}
