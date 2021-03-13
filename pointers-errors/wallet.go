package main

import (
	"errors"
	"fmt"
)

type CyberCurrency int

func (c CyberCurrency) String () string {
	return fmt.Sprintf("%d cC", c)
}

type Wallet struct {
	balance CyberCurrency
}

func (w *Wallet) Deposit (amount CyberCurrency) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw (amount CyberCurrency) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance () CyberCurrency {
	return w.balance
}
