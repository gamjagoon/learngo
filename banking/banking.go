package banking

import (
	"errors"
	"fmt"
)

// Account is good
type Account struct {
	owner   string
	balance int
}

// NewAccount is create strunt
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw your balance
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw you are poor")
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner return name
func (a *Account) Owner() string {
	return a.owner
}

func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
