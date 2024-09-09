package main

import "fmt"

func main() {
	var number [3]int
	number[0] = 42
	number[2] = 108
	var letters = [3]string{"a", "b", "c"}
	for index,value := range letters{
		fmt.Println(index,value)
	}

}
