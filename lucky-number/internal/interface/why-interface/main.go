package main

import (
	"fmt"
)

type Animal interface {
	Speak()
}


type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Cow struct {
	Name string
}

func (d Dog) Speak() { fmt.Println("ঘেউ ঘেউ") }
func (c Cat) Speak() { fmt.Println("মিউ মিউ") }
func (c Cow) Speak() { fmt.Println("হাম্বা") }

func main() {

	// if you don't use interface you can't do this.. you can say that using interface you can group different struct in a group 
	animals := []Animal{
		Dog{Name: "ঘেউ"},
		Cat{Name: "মিউ"},
		Cow{Name: "হাম্বা"},
	}
	// without using interface you can't write animal.Speak() 
	for _ , animal := range animals {
		animal.Speak()
	}

	

}
