package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const totalStudents = 200

func main() {
	rand.NewSource(time.Now().UnixNano())
	var wg sync.WaitGroup
	ratings := make([]int, totalStudents)

	for i := 0; i < totalStudents; i++ {
		wg.Add(1)

		go func(index int) {
			defer wg.Done()

			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

			rating := rand.Intn(10) + 1
			ratings[index] = rating
			fmt.Printf("Student %d gave a rating: %d\n", index+1, rating)
		}(i)
	}

	wg.Wait()

	total := 0
	for _, rating := range ratings {
		total += rating
	}
	averageRating := float64(total) / float64(totalStudents)

	fmt.Printf("\nAverage Teacher Rating: %.2f\n", averageRating)
}
