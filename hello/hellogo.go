package main

import (
	"fmt"

	"example.com/greetings"
)

var pl = fmt.Println

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Minh Huy")
	pl(message)
}
