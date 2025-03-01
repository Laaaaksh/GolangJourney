package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// BankAccount struct with balance and mutex
type BankAccount struct {
	balance int
	mutex   sync.Mutex
}

// Deposit function
func (acc *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()  // Mark goroutine as done
	acc.mutex.Lock() // Lock to ensure only one update at a time
	acc.balance += amount
	fmt.Printf("Deposited Rs.%d, New Balance: Rs.%d\n", amount, acc.balance)
	acc.mutex.Unlock() // Unlock after update
}

// Withdraw function
func (acc *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()  // Mark goroutine as done
	acc.mutex.Lock() // Lock to ensure only one update at a time
	if acc.balance >= amount {
		acc.balance -= amount
		fmt.Printf("Withdrawn Rs.%d, New Balance: Rs.%d\n", amount, acc.balance)
	} else {
		fmt.Printf("Failed to withdraw Rs.%d, Insufficient Balance: Rs.%d\n", amount, acc.balance)
	}
	acc.mutex.Unlock() // Unlock after update
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random generator
	var wg sync.WaitGroup
	account := &BankAccount{balance: 500} // Initialize account with Rs.500

	// Simulate multiple deposits and withdrawals concurrently
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go account.Deposit(rand.Intn(500)+1, &wg)  // Deposit random amount between 1-500
		go account.Withdraw(rand.Intn(500)+1, &wg) // Withdraw random amount between 1-500
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Final Account Balance: Rs.%d\n", account.balance)
}
