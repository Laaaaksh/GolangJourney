package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var mutex sync.Mutex

func Deposit(amount int, balance *int) {
	mutex.Lock()
	*balance += amount
	fmt.Println("Deposited ", amount)
	mutex.Unlock()
}
func Withdraw(amount int, balance *int) {
	mutex.Lock()
	if amount > *balance {
		fmt.Println("Insufficient balance")
	} else {
		fmt.Println("Withdrawn ", amount)
		*balance -= amount
	}

	mutex.Unlock()
}

func main() {

	balance := 500
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Deposit(rand.Intn(100)+1, &balance)
		}()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Withdraw(rand.Intn(200)+1, &balance)
		}()
	}

	wg.Wait()
	fmt.Println("Final balance is ", balance)
}
