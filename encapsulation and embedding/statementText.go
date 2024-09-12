package main

import (
	"errors"
	"fmt"
	"log"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

func (c *Coordinates) Latitude() float64 {
	return c.latitude
}
func (c *Coordinates) Longitude() float64 {
	return c.longitude
}
func (c *Coordinates) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude")
	}
	c.latitude = latitude
	return nil
}
func (c *Coordinates) SetLongtitude(longtitude float64) error {
	if longtitude < -180 || longtitude > 180 {
		return errors.New("invalid longtitude")
	}
	c.longitude = longtitude
	return nil
}

func main() {
	coordinates := Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongtitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())

}
