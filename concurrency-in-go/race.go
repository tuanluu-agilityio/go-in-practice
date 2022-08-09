package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// Track words in a struct
type words struct {
	found map[string]int
}

// Tracks the number of times you've seen this word
func (w *words) add(word string, n int) {
	// If the word isn't already tracked, add it.
	// Otherwise, increment the count.
	count, ok := w.found[word]
	if !ok {
		w.found[word] = n
		return
	}
	w.found[word] = count + n
}

// Create a new words instance
func newWords() *words {
	return &words{
		found: map[string]int{},
	}
}

/*
Open a file, parse its contents, and count the words that appear.
Copy function does all the copying for you.
*/
func tallyWords(filename string, dict *words) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Scanner is a useful tool for parsing files like this.
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict.add(word, 1)
	}
	return scanner.Err()
}

func main() {
	// Use a wait group to monitor a group of goroutines
	var wg sync.WaitGroup

	w := newWords()
	for _, f := range os.Args[1:] {
		wg.Add(1)
		go func(file string) {
			if err := tallyWords(file, w); err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	// Print what you found
	fmt.Println("Words that appear more than once:")
	for word, count := range w.found {
		if count > 1 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
}