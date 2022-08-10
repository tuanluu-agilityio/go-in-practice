package main

import "time"

func main() {
	ch := make(chan bool)
	timeout := time.After(600 * time.Millisecond)
	go send(ch)
	// Loops over a select with two channels and a default
	for {
		select {
		// If you get a message over your main channel, prints something
		case <-ch:
			println("Got message")
		// If a time-out occurs, terminates the program
		case <-timeout:
			println("Time out")
		// By default, sleeps for a bit. This makes the example easier to work with.
		default:
			println("*yawn*")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Sends a single message over the channel and then closes the channel
func send(ch chan bool) {
	time.Sleep(120 * time.Millisecond)
	ch <- true
	close(ch)
	println("Send and closed")
}