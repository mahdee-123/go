package main
// standarr libary
import (
	"fmt"
	"time"
)

type BankAccount struct {
	AccountNumber int
	Balance       float64
}

func Deposit(acc *BankAccount, amount float64) {
	acc.Balance += amount
	time.Sleep(time.Second*5)
}
func main() {

	acc := BankAccount{AccountNumber: 11023, Balance: 100.0}
	fmt.Println(acc)
	go Deposit(&acc, 1200)
	go Deposit(&acc, 1200)
	fmt.Println(acc.Balance)
	fmt.Println(acc.Balance)

}