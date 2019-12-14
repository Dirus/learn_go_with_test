package pointerserrors

import (
	"errors"
	"fmt"
)

// Bitcoin ...
type Bitcoin int

// Stringer ...
type Stringer interface {
	String() string
}

// String ...
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

// Deposit ...
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds ...
var ErrInsufficientFunds = errors.New("cannot wothdraw, insufficient funds")

// Withdraw ...
func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= b
	return nil
}
