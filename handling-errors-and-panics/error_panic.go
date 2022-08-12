package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("Can't divide by zero")

func main() {
	fmt.Println("Divide 1 by 0")
	// First you divide using the precheckDivide function, which returns an error.
	_, err := precheckDivide(1, 0)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	// Then you run a similar division, but with the divide function.
	fmt.Println("Divide 2 by 0")
	divide(2, 0)
}

// Returns an error if the divisor is 0.
func precheckDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return divide(a, b), nil
}

// The regular divide function wraps the division operator with no checks.
func divide(a, b int) int {
	return a / b
}