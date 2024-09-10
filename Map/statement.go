package main

import "fmt"

func main() {
	var myMap map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	var ranks = make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(myMap)
	fmt.Println(ranks)

}
