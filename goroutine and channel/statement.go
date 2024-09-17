package main

import (
	"fmt"
	"time"
)

func a() {
	for i := 1; i <= 50; i++ {
		fmt.Print("a")
	}
	fmt.Println()
}
func b() {
	for i := 1; i <= 50; i++ {
		fmt.Print("b")
	}
	fmt.Println()
}
func main() {
	go a()
	go b()
	time.Sleep(time.Second)
	fmt.Println("end main()")

}
