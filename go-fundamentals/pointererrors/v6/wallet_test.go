package v6

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		// run go linter called errcheck in terminal to check errors that missing error handling
		// to install: go install github.com/kisielk/errcheck@latest
		// to use it inside a folder: errcheck .
		// errcheck is a program for checking for unchecked errors in Go code.
		err := wallet.WithDraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		// in Go to indicate an error it is idiomatic to function return an err to check and act on.
		err := wallet.WithDraw(Bitcoin(15))
		assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()
	// nil is synonymous with null from other programming languages.
	// and it can be nil because it's an error interface.
	// interface in Go can be nil
	if got == nil {
		// fatal will stop the test if it's called.
		t.Fatal("wanted an error but didnt' get one")
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got.Error(), want)
	}
}
