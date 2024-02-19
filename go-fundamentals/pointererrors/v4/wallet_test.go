package v4

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.WithDraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		// in Go to indicate an error it is idiomatic to function return an err to check and act on.
		err := wallet.WithDraw(Bitcoin(15))
		assertBalance(t, wallet, startingBalance)

		// nil is synonymous with null from other programming languages.
		// and it can be nil because it's an error interface.
		// interface in Go can be nil
		if err == nil {
			t.Error("wanted an error but didnt' get one")
		}
	})
}
