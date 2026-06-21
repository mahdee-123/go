package main

import "fmt"

type Animaal interface {
	speak()
}

type Dog struct {
	name string
}

func (d Dog) speak() {
	fmt.Println("vew vew,,,,")
}
func main() {

	var animal Animaal = Dog{"dog"}
	// fmt.Println(animal.name) // we can't acces name

	dog , ok := animal.(Dog) 
	if ok {
		fmt.Println(dog.name) 
	} else {
		fmt.Println("invalid")
	}

}