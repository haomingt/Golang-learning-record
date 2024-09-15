package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func OpenFile(fileName string) (*os.File,error) {
	fmt.Println("Opening",fileName)
	return os.Open(fileName)
}
func CloseFile(file *os.File) {
	fmt.Println("Closing file")
	file.Close()
}
func GetFloats(fileName string) ([]float64,error) {
	var number []float64
	file,err := OpenFile(fileName)
	if err != nil {
		return nil,err
	}
	defer CloseFile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str,err := strconv.ParseFloat(scanner.Text(),64)
		if err != nil {
			return nil,err
		}
		number = append(number, str)
	}
	
	if scanner.Err() != nil {
		return nil,err
	}
	return number,nil
} 
func main() {
	numbers,err := GetFloats(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _,val := range numbers {
		sum += val
	}
	fmt.Printf("%0.2f",sum)
}