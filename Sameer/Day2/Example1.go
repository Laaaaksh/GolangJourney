package main

import "fmt"

func main() {

	fruitsSlice := []string{"apple, banana, orange, mango, pineapple"}

	countMap := make(map[int]int)

	for i := 0; i < len(fruitsSlice); i++ {
		fmt.Println(fruitsSlice[i])

		for j := 0; j < len(fruitsSlice[i]); j++ {
			countMap[int(fruitsSlice[i][j])]++

		}
	}
	for i := 0; i < 26; i++ {

		fmt.Printf("Frequency of %c is ", rune(i+97))
		fmt.Print(countMap[i+97])
		fmt.Println()

	}

}
