package main

import "errors"

// Textio needs a new way to update user's account balance.

// Assignment
// Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error. Depending on the transactionType, it should either add or subtract the amount from the customer's balance. If the customer does not have enough money, it should return the error insufficient funds and leave the balance unchanged. If the transactionType isn't a withdrawal or deposit, it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil.

// alice := customer{id: 1, balance: 100.0}
// deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}

// updateBalance(&alice, deposit)
// // id: 1 balance: 150



type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(customer *customer, transaction transaction ) error {

	if transaction.amount > customer.balance {
		 return errors.New("insufficient funds")
	} 
		
	if transaction.transactionType == transactionDeposit {
		customer.balance += transaction.amount
	} else if transaction.transactionType == transactionWithdrawal {
		customer.balance -= transaction.amount
	} else {
		return errors.New("unknown transaction type")
	}
	return nil
}

func main() {

	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}
	updateBalance(&alice, deposit)
}
