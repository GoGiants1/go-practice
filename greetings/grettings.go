package greetings

import (
	"errors"

	"fmt"
)

// Hello returns a greetings for the named person.

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If name was received
	// Return a greeting that embeds the name in a msg.
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg, nil
}
