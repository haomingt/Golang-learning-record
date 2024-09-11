package main

import "fmt"

type MyType string
type Number int

func (m MyType) sayHi() {
	fmt.Println("Hi", m)
}
func (n *Number) Double() {
	*n *= 2
}
func main() {
	value := MyType("a MyType value")
	value.sayHi()
	val := Number(3)
	val.Double()
	fmt.Println(val)

}
