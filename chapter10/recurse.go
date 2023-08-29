package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"strconv"
)

func scanDirectory(path string) error {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			err := scanDirectory(filePath)
			if err != nil {
				fmt.Printf("Returning error from scanDirectory(\"%s\") call\n", path)
				panic(err)
			}
		} else {
			fmt.Println(filePath)
		}
	}
	return nil
}

func calmDown() {
	fmt.Println("calmDown")
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

type awardPrizeError int

func (a awardPrizeError) Error() string {
	fmt.Println("awardPrizeError Error")
	errstr := strconv.Itoa(int(a))
	return errstr + " + some text"

}

func awardPrize() {
	fmt.Println("awardPrize call")

	defer calmDown()
	doorNumber := rand.Intn(3) + 10
	if doorNumber == 1 {

		fmt.Println("You win a cruise!")
	} else if doorNumber == 2 {

		fmt.Println("You win a car!")
	} else if doorNumber == 3 {

		fmt.Println("You win a goat!")
	} else {
		panic(awardPrizeError(doorNumber))
	}
}

func main() {
	defer calmDown()

	defer awardPrize()
	err := scanDirectory("bad")
	fmt.Println("after scanDirectory")

	if err != nil {
		log.Fatal(err)
	}
}
