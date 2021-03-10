package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	want := CyberCurrency(10)

	wallet.Deposit(want)

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %d", got, want)
	}
}

