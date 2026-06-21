package main

import (
	"fmt"
)

func riskyOperation() { 
	
	
	defer func() {
		if r := recover();  r != nil {
			fmt.Println("recovered from panic..")
		}
	}() 
	fmt.Println("starting risky operation....")
	panic("something went wrong..")
	// fmt.Println("this won't run..")
}

func safeDivide(a,b int) (result int, err error) {

	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("can't divide by zero...")
		}
	}()		

	result = a / b
	return result, nil
}
func main() {
	// when a function calls panic, the programm crashes and prints a stack trace

	riskyOperation()
	safeDivide(1,0)

	fmt.Println("programm is going...")
}