package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const totalStudents = 200

func getRating(studentID int, ratingsChan chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when function exits

	rand.Seed(time.Now().UnixNano() + int64(studentID))          // Seed random number generator
	rating := rand.Intn(10) + 1                                  // Generate a random rating between 1 and 10
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) // Simulate delay

	ratingsChan <- rating // Send rating to channel
}

func main() {
	var wg sync.WaitGroup
	ratingsChan := make(chan int, totalStudents) // Buffered channel to store ratings

	// Spawn goroutines for each student
	for i := 1; i <= totalStudents; i++ {
		wg.Add(1)
		go getRating(i, ratingsChan, &wg)
	}

	// Close the channel once all ratings are collected
	go func() {
		wg.Wait()
		close(ratingsChan)
	}()

	// Collect and compute the average rating
	totalRating := 0
	count := 0
	for rating := range ratingsChan {
		totalRating += rating
		count++
	}

	averageRating := float64(totalRating) / float64(count)
	fmt.Printf("Average Teacher Rating: %.2f\n", averageRating)
}
