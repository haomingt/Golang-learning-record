package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}
func sayBye() {
	fmt.Println("Bye")
}
func twice(theFunction func()) {
	theFunction()
	theFunction()
}
func main() {
	var myFunction func()
	myFunction = sayHi
	myFunction()
	twice(sayHi)
	twice(sayBye)

}
