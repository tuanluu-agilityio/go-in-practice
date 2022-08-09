package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	// A WaitGroup doesn't need to be initialized.
	var wg sync.WaitGroup
	var i int = -1
	var file string

	// Collects a list of files passed in on the command line
	for i, file = range os.Args[1:] {
		fmt.Println("File: ", file)
		// For every file you add, you tell the wait gorup that you're waiting for one
		// more compress operation.
		wg.Add(1)
		// This function calls compress and then notifies the wait group that it's done.
		go func(filename string) {
			compress(filename)
			wg.Done()
		}(file)
	}
	// The outer goroutine(main) waits until all the compressing goroutines have called wg.Done.
	wg.Wait()

	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	// Opens the source file for reading
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	// Opens a destination file, with the .gz extension added to the source file's name
	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// The gzip.Writer compresses data and then write it to the underlying file.
	gzout := gzip.NewWriter(out)

	// The io.Copy function does all the copying for you.
	_, err = io.Copy(gzout, in)
	gzout.Close()

	return err
}