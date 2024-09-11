package main

import "fmt"

type Gallons1 float64
type Liters1 float64
type Milliliters float64

func ToGallons(l Liters1) Gallons1 {
	return Gallons1(l * 0.264)
}
func ToLiters(g Gallons1) Liters1 {
	return Liters1(g * 3.785)
}
func (l Liters1) ToGallons() Gallons1 {
	return Gallons1(l * 0.264)
}
func (m Milliliters) ToGallons() Gallons1 {
	return Gallons1(m * 0.000264)
}
func main() {
	carFuel := Gallons1(1.2)
	busFuel := Liters1(4.5)
	carFuel += Liters1(40).ToGallons()
	fmt.Println(carFuel, busFuel)

}
