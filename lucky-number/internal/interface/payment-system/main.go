package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) 
}

type Bkashpaymnent struct{}
type Nagadpaymnent struct{}
type CreditCard struct{}
func (b Bkashpaymnent) Pay(amount float64) {
	fmt.Printf("bKash দিয়ে %0.2f টাকা পেমেন্ট হয়েছে\n", amount)
}

func (b Nagadpaymnent) Pay(amount float64) {
	fmt.Printf("Nagad দিয়ে %0.2f টাকা পেমেন্ট হয়েছে\n", amount)
}

func (b CreditCard) Pay(amount float64) {
	fmt.Printf("CreditCard দিয়ে %0.2f টাকা পেমেন্ট হয়েছে\n", amount)
}


func ProcessOrder(p PaymentMethod, amount float64) {
	p.Pay(amount)
	fmt.Println("অর্ডার সম্পন্ন!")
}

func main() {
    ProcessOrder(CreditCard{},500)
}