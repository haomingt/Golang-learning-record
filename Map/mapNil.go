package main

import "fmt"

func main() {
	counters := map[string]int{"a": 0, "b": 2}
	value, ok := counters["a"]
	fmt.Println(value,ok)
	value,ok = counters["b"]
	fmt.Println(value,ok)
	value,ok = counters["c"]
	fmt.Println(value,ok)
	fmt.Println(counters)
	
}