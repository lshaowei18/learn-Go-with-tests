package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	Balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.Balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.Balance < amount {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.Balance -= amount
	return nil
}

func (w *Wallet) CheckBalance() Bitcoin {
	return w.Balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.3f BTC", b)
}
