package main

import "fmt"

type Address struct {
	Street    string
	City      string
	State     string
	PotalCode string
}
type mycar struct {
	name     string
	topSpeed int
	Address
}
type Employee struct {
	Name   string
	Salary float64
	Address
}

func main() {
	address := Address{Street: "123", City: "45", State: "67", PotalCode: "56"}
	var e Employee
	e.Address = address
	fmt.Printf("%#v\n", e)
	fmt.Println(e.Address.City)
	fmt.Printf("%v\n", e.PotalCode)
}
