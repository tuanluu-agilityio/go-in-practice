package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Creates a channel that will receive a message when 30 seconds have elapsed
	done := time.After(30 * time.Second)
	// Make a new channel for passing bytes from Stdin to Stdout.
	echo := make(chan []byte)
	// Starts a goroutine to read Stdin, passes it our new channel for communicating
	go readStdin(echo)
	for {
		// Uses a select statement to pass data from Stdin to Stdout when received,
		// or to shut down when the time-out event occurs
		select {
		case buf := <-echo:
			os.Stdout.Write(buf)
		case <-done:
			fmt.Println("Timed out")
			os.Exit(0)
		}
	}
}

// Takes a write-only channel(chan<-) and sends any received input to that channel
func readStdin(out chan<- []byte) {
	for {
		// Copies some data from Stdin into data.
		// Note that File.Read blocks until it receives data.
		data := make([]byte, 1024)
		l, _ := os.Stdin.Read(data)
		if l > 0 {
			// Sends the buffered data over the channel
			out <- data
		}
	}
}