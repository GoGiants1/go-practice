package greetings

import "fmt"

// Hello returns a greetings for the named person.

func Hello(name string) string {
	// Return a greeting that embeds the name in a msg.
	msg := fmt.Sprintf("Hi, %v. Welcome!", name)
	return msg
}
