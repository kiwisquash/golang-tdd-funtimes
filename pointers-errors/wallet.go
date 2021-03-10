package main

import "fmt"

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

func (w *Wallet) Balance () CyberCurrency {
	return w.balance
}
