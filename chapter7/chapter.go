package main

import (
	"fmt"
	"sort"
)

func exapmlePage252() {
	var ok bool
	ranks := make(map[string]int)
	var rank int
	ranks["bronze"] = 3
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	delete(ranks, "bronze")
	rank, ok = ranks["bronze"]
	fmt.Printf("rank: %d, ok: %v\n", rank, ok)
	isPrime := make(map[int]bool)
	var prime bool
	isPrime[5] = true
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
	delete(isPrime, 5)
	prime, ok = isPrime[5]
	fmt.Printf("prime: %v, ok: %v\n", prime, ok)
}

func taskPage251() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {

			fmt.Printf("%s: not found\n", letter)
		} else {

			fmt.Printf("%s: %d\n", letter, count)
		}
	}
}

func exapmlePage256() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	for name, grade := range grades {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grade)
	}
	fmt.Println("Class roster:")
	for name := range grades {
		fmt.Println(name)
	}
	fmt.Println("Grades:")
	for _, grade := range grades {
		fmt.Println(grade)
	}
}

func exapmlePage257() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var names []string
	for name := range grades {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}

func taskPage260() {
	ranks := map[string]int{"bronze": 3, "silver": 2, "gold": 1}
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", medal, rank)
	}
}

func main() {
	taskPage251()
	fmt.Printf("--------------------\n")
	exapmlePage252()
	fmt.Printf("--------------------\n")
	exapmlePage256()
	fmt.Printf("--------------------\n")
	exapmlePage257()
	fmt.Printf("--------------------\n")
	taskPage260()
}
