package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Deposit method for depositing money into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance method to check balance of wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Withdraw method to withdraw money from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
