package main

import "fmt"

type Truck1 string

func (t Truck1) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck1) Brake() {
	fmt.Println("Stopping")
}
func (t Truck1) Steer(dir string) {
	fmt.Println("Turning", dir)
}
func (t Truck1) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle1 interface {
	Accelerate()
	Brake()
	Steer(string)
}

func TryVehicle(vehicle Vehicle1) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck1)
	if ok {
		truck.LoadCargo("test cargo")
	} else {
		fmt.Println("test cargo")
	}
}
func main() {
	TryVehicle(Truck1("Fnord F180"))

}
