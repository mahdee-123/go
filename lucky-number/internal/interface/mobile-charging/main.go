package main

import "fmt"


type Chargeable interface {
	Charge() 
}

type Iphone struct{}
type Android struct{}

func (i Iphone) Charge() {
	fmt.Println("iphone is charging....")
}

func (i Android) Charge() {
	fmt.Println("Android is charging....")
}


func ChargeDevice(device Chargeable) {
	device.Charge()
}


func main() {
	// i := Iphone{}
	// i.Charge() 

	// a := Android{}
	// a.Charge() 


}