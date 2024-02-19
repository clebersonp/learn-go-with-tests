package v4

import "fmt"

// Bitcoin is a custom type from existing int type.
// It's more descriptive
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

func (w *Wallet) WithDraw(amount Bitcoin) {
	w.balance -= amount
}
