package main

import (
	"errors"
	"fmt"
	"log"
	"unicode/utf8"
)

type Date struct {
	year  int
	mouth int
	day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}
func (d *Date) SetMouth(mouth int) error {
	if mouth < 1 || mouth > 12 {
		return errors.New("invalid mouth")
	}
	d.mouth = mouth
	return nil
}
func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}
func (d *Date) Mouth() int {
	return d.mouth
}
func (d *Date) Day() int {
	return d.day
}
type Event struct {
	title string
	Date
}
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
func (e *Event) Title()string {
	return e.title
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
	err = date.SetDay(12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Day())
	fmt.Println(date.Mouth())
	fmt.Println(date.Year())
	event := Event{}
	err = event.SetTitle("kjasgjdsjksagdjdgjhasgdjashasdhasjgdggggggggggggggggggggggg")
	if err != nil {
		log.Fatal(err)
	}

}
