package main

import (
	"fmt"
	"sync"
)

// Function to count letter frequencies in a string and send them via a channel
func countLetters(word string, ch chan map[rune]int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the WaitGroup counter when the function completes
	freqMap := make(map[rune]int)
	for _, char := range word {
		freqMap[char]++
	}
	ch <- freqMap // Send the frequency map through the channel
}

func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	ch := make(chan map[rune]int, len(words)) // Buffered channel to collect frequency maps
	var wg sync.WaitGroup

	// a goroutine for each word
	for _, word := range words {
		wg.Add(1)
		go countLetters(word, ch, &wg)
	}

	// Close the channel once all goroutines complete
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Aggregate results from the channel
	totalFreq := make(map[rune]int)
	for freqMap := range ch {
		for char, count := range freqMap {
			totalFreq[char] += count
		}
	}

	// Print the aggregated frequency map
	fmt.Println("Letter Frequency:")
	for char, count := range totalFreq {
		fmt.Printf("%q: %d\n", char, count)
	}
}
