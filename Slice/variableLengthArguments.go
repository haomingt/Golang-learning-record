package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 {
	var max float64 = math.Inf(-1)
	for _, val := range numbers {
		max = math.Max(val, max)
	}
	return max
}
func inRange(minval float64, maxval float64, numbers ...float64) []float64 {
	var result = []float64{}
	for _, val := range numbers {
		if val >= minval && val <= maxval {
			result = append(result, val)
		}
	}
	return result
}
func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, val := range numbers {
		sum += val
	}
	return sum / float64(len(numbers))
}
func main() {
	fmt.Println(maximum(89.45, 89374, 75, 87, 86, 34, 56))
	ans := inRange(1, 100, 89, 78, 9289283, 2663, 28938, 89238, 77, 98)
	fmt.Println(ans)
	fmt.Println(average(6, 7, 6, 7,8))
	var numbers = []float64{4,5,6}
	fmt.Println(average(numbers...))
}
