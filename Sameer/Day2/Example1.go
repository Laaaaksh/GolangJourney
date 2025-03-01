package main

import (
	"fmt"
	"sync"
)

func countFreq(s string, ch chan map[int]int) {
	lettercount := make(map[int]int)
	for i := 0; i < len(s); i++ {

		lettercount[int(s[i])]++
	}
	ch <- lettercount
}

func main() {

	fruitsSlice := []string{"apple", "banana", "orange", "mango", "pineapple"}
	final := make(map[int]int)
	ch := make(chan map[int]int, len(fruitsSlice))

	var wg sync.WaitGroup

	for i := 0; i < len(fruitsSlice); i++ {
		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()
			countFreq(fruitsSlice[i], ch)
		}(i, fruitsSlice[i])
	}

	wg.Wait()
	close(ch)

	for i := 0; i < len(fruitsSlice); i++ {
		freqMap := <-ch
		for k, v := range freqMap {
			final[k] += v
		}
	}

	for i := 97; i < 122; i++ {
		fmt.Printf("%c : %d\n", i, final[i])
	}

}
