package main

import (
	"fmt"
	"errors"
)

func main() {
	defer func() {
		// Provides a deferred closure to handle panic recovery
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	// Call a functions that panics
	yikes()
}

func yikes() {
	// Emits a panic with an error for a body
	panic(errors.New("Something bad happened."))
}
