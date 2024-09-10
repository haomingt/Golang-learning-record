package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var mySlice = []float64{}
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str,err := strconv.ParseFloat(scanner.Text(),64)
		if err != nil {
			log.Fatal(err)
		}
		mySlice = append(mySlice, str)
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
	//计算平均值
	var sum float64 = 0
	for _,val := range mySlice {
		sum += val
	}
	fmt.Println("平均值为",sum / float64(len(mySlice)))
}
