package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	success := false
	fmt.Println("I've chosem a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make a guess:")
	for i := 10; i > 0; i-- {
		fmt.Println("you have", i, "guesses left")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Oops.Your guess was high")
		} else if guess < target {
			fmt.Println("Oops.Your guess was low")
		} else {
			fmt.Println("Good job!Your guessed it")
			success = true
			break
		}
	}
	if success == false {
		fmt.Println("Sorry,you didn't guess my number.It was:", target)
	}
}
