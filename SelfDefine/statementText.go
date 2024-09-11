package main

import "fmt"

type Population int

func main() {
	var population Population
	population = Population(572)
	fmt.Println(population)
	population += 1
	fmt.Println(population)

}
