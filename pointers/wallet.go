package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Go lets you create new types from existing ones. It's called underlying type
// type declaration
// Useful for adding more domain specific meaning to values
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// satisfying Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Technically you do not need to change Balance to use a pointer receiver as taking a copy of the balance is fine.
// However, by convention you should keep your method receiver types the same for consistency.
// *Wallet - a pointer to a wallet
func (w *Wallet) Balance() Bitcoin {
	// fmt.Printf("address of balance in wallet is %v \n", &w.balance)
	// equivalent: (*w).balance
	// language permits us to write w.balance, without an explicit dereference
	// These pointers to structs even have their own name: struct pointers and they are automatically dereferenced
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in wallet Deposit is %v\n", &w.balance)
	w.balance += amount
}

// Go copies values when you pass them to functions/methods,
// so if you're writing a function that needs to mutate state you'll need it
// to take a pointer to the thing you want to change
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
