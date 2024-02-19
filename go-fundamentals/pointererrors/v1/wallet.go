package v1

import "fmt"

type Wallet struct {
	// with lower case we are restricting its visibility, so balance property is private outside the package
	// to access its value we use Balance method to do that
	balance int
}

func (w Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in Deposit method is %p \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() int {
	return w.balance
}
