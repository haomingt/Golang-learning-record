package main

import "fmt"

func snack() {
	defer fmt.Println("C")
	fmt.Println("O")
	panic("panic")
}

func main() {
	snack()

}
