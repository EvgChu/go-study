package main

import (
	"fmt"
	"hfg/chapter5/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("chapter7/votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
