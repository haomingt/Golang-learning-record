package main

import (
	"fmt"
)


type ComedyError string

func (c ComedyError) Error() string {
	return string(c)
	
}

func main() {
	var err error
	err = ComedyError("What's a programmer's favorite beer?Logger!")
	fmt.Println(err)

}
