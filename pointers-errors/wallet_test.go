package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want CyberCurrency) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %d", got, want)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("expected an error.")
		}
	}

	t.Run("testing deposit", func(t *testing.T){

		wallet := Wallet{}
		want := CyberCurrency(10)

		wallet.Deposit(want)

		assertBalance(t, wallet, want)
	})

	t.Run("testing withdraw less than balance", func(t *testing.T){

		wallet := Wallet{}
		wallet.Deposit(CyberCurrency(17))
		wallet.Withdraw(CyberCurrency(10))

		want := CyberCurrency(7)

		assertBalance(t, wallet, want)

	})

	t.Run("testing withdraw over the balance", func(t *testing.T){

		wallet := Wallet{}
		startingBalance := CyberCurrency(17)
		wallet.Deposit(startingBalance)
		err := wallet.Withdraw(CyberCurrency(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})

}

