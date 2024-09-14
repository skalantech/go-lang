package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Peter", "Clark", "Bruce", "Neo", "Rick", "Darryl"}

	// Request greeting messages for the names.
	messages, err0 := greetings.Hellos(names)
	if err0 != nil {
		log.Fatal(err0)
	}
	// Request greeting message and print it.
	message, message2, err1 := greetings.Hello("John")
	// If an error was returned, print it to the console and
	// exit the program.
	if err1 != nil {
		log.Fatal(err1)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println(message2)
	fmt.Println(messages)
}
