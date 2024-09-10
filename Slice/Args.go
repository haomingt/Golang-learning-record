package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	os.Args = os.Args[1:]
	var sum float64 = 0
	for _, val := range os.Args {
		number, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Printf("平均值为%f", sum/float64(len(os.Args)))

}
