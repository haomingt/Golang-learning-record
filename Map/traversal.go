package main

import "fmt"

func main() {
	ranks := make(map[string]int)
	ranks["a"] = 1
	ranks["b"] = 2
	ranks["c"] = 3
	for key, value := range ranks {
		fmt.Println("键值为：", key, "所对应的值为：", value)
	}
}
