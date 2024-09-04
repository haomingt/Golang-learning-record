package main

import (
	"fmt"
)

func double1(num *int) {
	*num *= 2
}
func main() {

	amount := 2
	double1(&amount)
	fmt.Println(amount)
}
