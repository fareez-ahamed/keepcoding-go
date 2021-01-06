package main

import (
	"errors"
	"fmt"
	"sync"
)

// Account is a bank account type
type Account struct {
	balance int
	sync.RWMutex
}

// Balance is to check the account balance
func (a *Account) Balance() int {
	a.RLock()
	defer a.RUnlock()
	return a.balance
}

// Deposit is to add money
func (a *Account) Deposit(amount int) {
	a.Lock()
	a.balance = a.balance + amount
	a.Unlock()
}

// Withdraw is to take out money
func (a *Account) Withdraw(amount int) error {
	a.Lock()
	defer a.Unlock()
	if a.balance >= amount {
		a.balance = a.balance - amount
		return nil
	}
	return errors.New("Account balance not sufficient")
}

func main() {
	acc := Account{balance: 500}
	done := make(chan struct{})

	go func() {
		defer func() { done <- struct{}{} }()
		err := acc.Withdraw(200)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Successfully withdrawn 200")
	}()

	go func() {
		defer func() { done <- struct{}{} }()
		err := acc.Withdraw(500)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Successfully withdrawn 500")
	}()

	<-done
	<-done
}
