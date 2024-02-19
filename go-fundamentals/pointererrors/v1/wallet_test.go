package v1

import "testing"

func TestWallet(t *testing.T) {
	// create an empty wallet
	wallet := Wallet{}

	// put amount into wallet
	wallet.Deposit(10)

	// return the current wallet amount
	got := wallet.Balance()
	want := 10

	// assertion
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
