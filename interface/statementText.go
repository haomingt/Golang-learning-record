package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(dir string) {
	fmt.Println("Turning", dir)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(dir string) {
	fmt.Println("Turning", dir)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(dir string)
}

func main() {
	var vehicle Vehicle = Car("Toyoda Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle = Truck("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
}
