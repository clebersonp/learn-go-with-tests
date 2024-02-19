package v6

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds use the 'var' keyword. It allows us to define values global to the package level
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) WithDraw(amount Bitcoin) error {
	if amount > w.balance {
		// errors.New creates a new error with a custom message
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
