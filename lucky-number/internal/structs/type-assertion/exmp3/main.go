package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

func (c Cat) MakeSound() string {
	return "meow"
}

func (d Dog) MakeSound() string {
	return "woof"
}
func main() {
	animals := []Animal{Dog{Name: "Dog"}, Cat{Name: "Cat"}}

	for _, animal := range animals {
		if dog, ok := animal.(Dog); ok {
			fmt.Println(dog.Name)
		} else if cat, ok := animal.(Cat); ok {
			fmt.Println(cat.Name)
		}
	}
}