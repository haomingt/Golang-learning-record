package main

import (
	"fmt"
)

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.active = true
	s.rate = 4.99
	return &s
}
func printInfo(s *subscriber) {
	fmt.Println(s.name)
	fmt.Println(s.rate)
	fmt.Println(s.active)
}
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
func main() {
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
	s1 := defaultSubscriber("chm")
	printInfo(s1)

}
