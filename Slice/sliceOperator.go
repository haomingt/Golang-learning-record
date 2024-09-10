package main

import "fmt"

func main() {
	myArray := [5]string{"a", "b", "c", "d", "f"}
	mySlice := myArray[0:3]
	mySlice[0] = "y"
	fmt.Println(myArray)
}
