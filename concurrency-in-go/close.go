package main

import (
	"fmt"
	"time"
)

/*
This example demonstrates a pattern that you'll frequently observe in Go: using a
channel (oftern called done) to send a signal between goroutines. In this pattern, you
usually have one goroutine whose primary task is to receive messages, and another
whose job is to send messages. If the receiver hits a stopping condition, it must let the
sender know.

The main function is the one that knows when to stop processing.
But it's also the receiver. The receiver shouldn't ever close a receiving channel. Instead,
it sends a message on the done channel indicating that it's done with its work. Now, the
send function knows when it receives a message on done that it can (and should) close the
channel and return.
*/
func main() {
	msg := make(chan string)
	// Adds an additional Boolean channel that indicates when you're finish
	done := make(chan bool)
	until := time.After(5 * time.Second)

	// Passes two channels into send
	go send(msg, done)
	
	for {
		select {
		case m := <-msg:
			fmt.Println(m)
		case <-until:
			// When you time-out, lets send know the process is done
			done <- true
			time.Sleep(500 * time.Millisecond)
			return
		}
	}
}

// ch is a receiving channel, while done is a sending channel.
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		// When done has message, shuts things down
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}