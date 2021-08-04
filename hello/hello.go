package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println("--------------")
	message := greetings.Hello("Gloria")
	fmt.Println(message)
}
