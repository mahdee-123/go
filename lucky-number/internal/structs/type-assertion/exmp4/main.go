package main

import (
	"errors"
	"fmt"
)

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64)   error
}

type CheckingAccount struct {
	Balance float64
}

func (acc *CheckingAccount) Deposit(amount float64) {
	acc.Balance += amount
}

func (acc *CheckingAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if amount > acc.Balance {
		return errors.New("insufficient balance")
	} 
	acc.Balance -= amount

	return nil
}

func main() {
	var acc BankAccount
	acc = &CheckingAccount{Balance: 100.0}

	acc.Deposit(60.0)
	fmt.Println(acc.Withdraw(50))
	myacc := acc.(*CheckingAccount)
	fmt.Println(myacc.Balance)
}