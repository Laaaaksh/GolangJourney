package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func takeRating() int {

	randomNumber := rand.Intn(5) + 1
	rating := rand.Intn(10) + 1
	time.Sleep(time.Duration(randomNumber) * time.Second)

	return rating

}

func main() {

	rating := 0

	var wg sync.WaitGroup

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rating += takeRating()
		}()
	}
	wg.Wait()
	fmt.Println("Average rating is ", rating/200)

}
