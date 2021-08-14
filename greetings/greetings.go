package greetings

import "fmt"

// Hello returns a greeting for the named user
func Hello(name string) string {
	// Returns a greeting with the name provided
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
