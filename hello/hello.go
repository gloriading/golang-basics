package main

import (
	"fmt"

	"example.com/greetings"

	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(messages["Darrin"])

	// message, err := greetings.Hello("Gloria")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)
}
