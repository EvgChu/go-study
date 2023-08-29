package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
		return err
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				panic(err)
				fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
				return err
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func calmDown() {
	recover()
}

func awardPrize() {
	defer calmDown()
	doorNumber := rand.Intn(3) + 10
	if doorNumber == 1 {

		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {

		fmt.Println("You win a car!")
	} else if doorNumber == 3 {

		fmt.Println("You win a goat!")
	} else {
		panic(doorNumber)
		fmt.Println("invalid door number")

	}
}

func main() {
	defer awardPrize()
	defer calmDown()
	err := scanDirectory("/")
	if err != nil {
		log.Fatal(err)
	}
}
