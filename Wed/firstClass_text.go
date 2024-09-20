package main

import "fmt"

func callFuntion(passedFuntion func()) {
	passedFuntion()
}
func callTwice(passedFuntion func()) {
	passedFuntion()
	passedFuntion()
}
func callWithArguments(passedFuntion func(string, bool)) {
	passedFuntion("This sentence is", false)
}
func printReturnValue(passedFuntion func() string) {
	fmt.Println(passedFuntion())
}
func functionA() {
	fmt.Println("function called")
}
func functionB() string {
	fmt.Println("function called")
	return "Returning from function"
}
func functionC(a string, b bool) {
	fmt.Println("function called")
	fmt.Println(a, b)
}

func main() {
	callFuntion(functionA)
	callTwice(functionA)
	callWithArguments(functionC)
	printReturnValue(functionB)
}
