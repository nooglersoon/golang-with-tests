package pointerserrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

// Bitcoin representsThe var keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	// In Go if a symbol (variables, types, functions et al) starts with a lowercase symbol then it is private outside the package it's defined in.
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// In Go, when you call a function or a method the arguments are copied.
	// When calling func (w Wallet) Deposit(amount int) the w is a copy of whatever we called the method from.

	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	// Remember we can access the internal balance field in the struct using the "receiver" variable.
	return w.balance
}

// Add asterisk prefix to Wallet type. The difference is the receiver type is *Wallet rather than Wallet which you can read as "a pointer to a wallet".

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
