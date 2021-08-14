package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named user
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// Returns a greeting with the name provided
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
