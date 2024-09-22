package main

import "fmt"

func main() {
	var str  = []rune("我爱中国")
	for _, val := range str {
		fmt.Println(string(val))
	}

}