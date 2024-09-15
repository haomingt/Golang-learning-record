package main

func AcceptAnything(thing interface{}) {

}

func main() {
	AcceptAnything("true")
	AcceptAnything(true)
	AcceptAnything(23)
	AcceptAnything(2.33)

}
