package main

import (
	"fmt"
	"log"
)

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees", o)
}
func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}
func main() {
	err := checkTemperature(121, 90)
	if err != nil {
		log.Fatal(err)
	}

}
