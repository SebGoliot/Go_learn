package main

import (
	"fmt"

	"ktu/greetings"
)

func main() {
	// Gets a greeting message and prints in.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
