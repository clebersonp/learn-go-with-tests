package v2

import "fmt"

// Bitcoin is a custom type from existing int type.
// It's more descriptive
type Bitcoin int

type Wallet struct {
	// with lower case we are restricting its visibility, so balance property is private outside the package
	// to access its value we use Balance method to do that
	balance Bitcoin
}

// Deposit - putting the notation * in front of Wallet type we are saying to Go that
// this is special type - "a pointer (address memory) to a wallet
// Therefore, when we change this wallet value, we are changing the wallet value in memory
// to access the value of wallet we need to dereference the pointer like: (*w).balance, but the language permits us to
// write w.balance, without explicit dereference. Go will make that for us automatically
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit method is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
