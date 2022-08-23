package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	var file io.ReadCloser
	// Runs OpenCSV and handles any errors. This implementation
	// always return an error.
	file, err := OpenCSV("data.csv")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// Uses a deferred function to ensure that a file gest closed
	defer file.Close()
}

// OpenCSV opens and preprocesses your file.
// Note the named return values.
func OpenCSV(filename string) (file *os.File, err error) {
	// The main deferred error handling happens here.
	defer func() {
		if r := recover(); r != nil {
			file.Close()
			err = r.(error)
		}
	}()
	// Opens the data file and handles any errors (such as file not found)
	file, err = os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file\n")
		return file, err
	}
	// Runs our intentionally broken RemoveEmptyLines function
	RemoveEmptyLines(file)

	return file, err
}

func RemoveEmptyLines(f *os.File) {
	// Instead of stripping empty lines, you always fail here.
	panic(errors.New("Failed parse"))
}
