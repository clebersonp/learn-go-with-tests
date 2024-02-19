package v2

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	// create an empty wallet
	wallet := Wallet{}

	// put amount into wallet
	wallet.Deposit(10)

	// return the current wallet amount
	got := wallet.Balance()

	// In Go, when you call a function or a method the arguments are copied
	// When calling func (w Wallet) Deposit(amount int) the w is a copy this wallet
	// to access this wallet address in memory we use symbol '&' in front of a wallet.
	fmt.Printf("address of balance in test is %p \n", &wallet.balance)

	want := 10

	// assertion
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
