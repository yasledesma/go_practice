package wallet

import (
	"errors"
	"fmt"
)

var ErrorInsufficientFunds = errors.New("Insufficient funds to perform this transaction.")

type Shitcoin int

type Wallet struct {
    balance Shitcoin
}

func (s Shitcoin) String() string {
    return fmt.Sprintf("%d STC", s)
}

func (w *Wallet) Deposit(amount Shitcoin) {
    w.balance += amount
}

func (w *Wallet) Balance() Shitcoin {
    return w.balance
}

func (w *Wallet) Withdraw(amount Shitcoin) error {
    if amount > w.balance {
        return ErrorInsufficientFunds
    }

    w.balance -= amount
    return nil
}
