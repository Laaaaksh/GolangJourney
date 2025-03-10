package main

import (
    "fmt"
    "strings"
)

func countLetters(word string, ch chan map[rune]int) {
    letterCount := make(map[rune]int) 
    for _, char := range strings.ToLower(word) { 
        letterCount[char]++
    }
    ch <- letterCount
}

func main() {
    words := []string{"quick", "brown", "fox", "lazy", "dog"}
    result := make(map[rune]int) 
    ch := make(chan map[rune]int) 

    for _, word := range words {
        go countLetters(word, ch) 
    }

    for i := 0; i < len(words); i++ {
        letterMap := <-ch 
        for char, count := range letterMap {
            result[char] += count 
        }
    }

    fmt.Println("Letter frequencies:")
    for char, count := range result {
        fmt.Printf("%q: %d\n", char, count)
    }
}
