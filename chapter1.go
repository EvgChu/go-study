package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
