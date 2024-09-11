package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Mouth int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.Year = year
	return nil
}
func (d *Date) SetMouth(mouth int) error {
	if mouth < 1 || mouth > 12 {
		return errors.New("invalid mouth")
	}
	d.Mouth = mouth
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.Day = day
	return nil
}
func main() {
	// data := Date{Year: 2019, Mouth: 5, Day: 27}
	// fmt.Println(data)
	var date Date
	err := date.SetYear(2024)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMouth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v", date)
}
