package main

import (
	"hfg/chapter4/dates"
	"fmt"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days + dates.DaysInWeek, "days")
}