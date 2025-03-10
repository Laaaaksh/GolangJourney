package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (acc *BankAccount) Deposit(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	acc.mu.Lock()
	defer acc.mu.Unlock()

	acc.balance += amount
	fmt.Printf("Deposited Rs.%d | New Balance: Rs.%d\n", amount, acc.balance)
}

// Withdraw method (removes money, but only if sufficient balance exists)
func (acc *BankAccount) Withdraw(amount int, wg *sync.WaitGroup) {
	defer wg.Done()

	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.balance >= amount {
		acc.balance -= amount
		fmt.Printf("Withdrawn Rs.%d | New Balance: Rs.%d\n", amount, acc.balance)
	} else {
		fmt.Printf("Failed to Withdraw Rs.%d | Insufficient Balance: Rs.%d\n", amount, acc.balance)
	}
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	account := BankAccount{balance: 500}
	var wg sync.WaitGroup

	// Simulate concurrent deposits and withdrawals
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go account.Deposit(rand.Intn(200)+50, &wg)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go account.Withdraw(rand.Intn(200)+50, &wg)
	}

	wg.Wait() // Wait for all goroutines to complete
	fmt.Printf("\nFinal Account Balance: Rs.%d\n", account.balance)
}
