package pointers

import (
	"fmt"
	"errors"
)

type Stringer interface {
	Strint() string
}

type Bitcoin int

func (b Bitcoin) String() string { // this stringer method is used to define how will the type is printed when used with %s place holder
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(transaction Bitcoin){ // this is a pointer to a wallet.
	w.balance += transaction
}

func (w *Wallet) Withdraw(amount Bitcoin) error{ // We want withdraw to return an error if money is overdrawn
	if amount > w.balance{
		return errors.New("Cannot withdraw, insufficient balance")
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin{ // the w here is the copy of the one, which called it. So any change that is going to happen is going to be this one and not reflected in the one the called it.
	return w.balance
}