package main

import (
	"fmt"
	"hfg/chapter8/magazine"
)

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func examplePage271() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
}

func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func examplesPage272to273() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)
	subscriber2 := defaultSubscriber("Beth Ryan")
	applyDiscount(&subscriber2)
	printInfo(subscriber2)
}

func nitroBoost(c *car) {
	c.topSpeed += 50
}

func doublePack(p *part) {
	p.count *= 2
}

func examplePage280() {
	var fuses part
	fuses.description = "Fuses"
	fuses.count = 5
	doublePack(&fuses)
	fmt.Println(fuses.description)
	fmt.Println(fuses.count)
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topSpeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topSpeed)
}

func examplePage289() {
	address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address = address
	fmt.Println(subscriber.Address)
	fmt.Println(subscriber.Street)
}

func main() {
	examplePage271()
	examplesPage272to273()
	examplePage280()
	examplePage289()
}
