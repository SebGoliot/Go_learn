package main

import (
	"fmt"
	"log"

	"ktu/greetings"
)

func main() {
	// Sets the properties of the prefdefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request a greeting message.
	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console, and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returns, print the returned message
	// to the console
	fmt.Println(message)
}
