package main

import "fmt"

func main() {
	var myArray []int = []int{7,6,4,3,2}
	mySlice := make([]string, 7)
	fmt.Println(len(mySlice))
	fmt.Println(myArray)
}