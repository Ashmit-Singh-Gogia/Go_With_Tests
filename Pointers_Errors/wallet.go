package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("\nAddress of Wallet balance in the function is: %p\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("\nBalance address in func Balance is: %p\n", &w.balance)
	return w.balance
}

var ErrInsfficientFunds = "insufficient funds, Try Smaller amount"

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.Balance() {
		return errors.New(ErrInsfficientFunds)
	}
	w.balance -= amount
	return nil
}
