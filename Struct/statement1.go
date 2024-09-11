package main

import "fmt"

func main() {
	var mysTruct struct {
		number float64
		word   string
		toggle bool
	}
	fmt.Printf("%#v\n", mysTruct)
	mysTruct.number = 56.90
	mysTruct.word = "987"
	mysTruct.toggle = true
	fmt.Printf("%#v\n", mysTruct)
	
}
