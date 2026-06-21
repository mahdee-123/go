package main

import "fmt"

type Speaker interface {
	speak()
	walk()
}

type cat struct{}

func (c cat) speak() {}
func (c cat) walk()  {}

func main() {
	var s Speaker

	s = cat{}

	fmt.Print(s)
}