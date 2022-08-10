package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	until := time.After(5 * time.Second)

	// Starts a send goroutine with a sending channel
	go send(msg)

	// Loops over a select that watches for messages from send, or for a time-out
	for {
		select {
		// If a message arrives from send, prints it
		case m := <- msg:
			fmt.Println(m)
		// When the time-out occurs, shuts things down.
		case <-until:
			close(msg)
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

// Sends "hello" to the channel every half-second
func send(ch chan string) {
	for {
		ch <- "hello"
		time.Sleep(500 * time.Millisecond)
	}
}