package main

import (
	"fmt"

	"example.com/greetings"

	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message1, err := greetings.Hello("Gloria")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message1)
}
