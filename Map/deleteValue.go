package main

import "fmt"

func main() {
	ranks := make(map[string]int)
	ranks["a"] = 1
	ranks["b"] = 2
	delete(ranks, "a")
	_, ok := ranks["a"]
	fmt.Println(ok)
	_, ok = ranks["b"]
	fmt.Println(ok)

}
