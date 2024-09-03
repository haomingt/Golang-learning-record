package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := paintNeeded(2.034, -3.067)
	if err != nil {
		log.Fatal(err)
	}
}
func double(num float64) float64 {
	return num * 2
}
func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	aera := width * height / 10
	fmt.Printf("%.2f liters needed\n", aera)
	return aera, nil
}
