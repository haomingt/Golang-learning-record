package main

import "fmt"

func calmDown() {
	recover()
}
func freakOut() {
	defer calmDown()
	panic("oh,no")
	fmt.Println("I won't be run")
}

func main() {
	freakOut()
	fmt.Println("Exiting normally")
}
