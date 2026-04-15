package main

import (
	"fmt"

	"github.com/mahdee-123/go-two/utils"
)

// variables , type and operations
var name string = "mahdee"
var age int = 23
var isCool bool = true
var name2, age2, isCool2 = "arnob", 21, false
// infer data type
var id = 1110 

// constants	
const pi float32  = 3.14
const appName string = "pussho"
func main() {
	
	sum := utils.Add(23,4)
	fmt.Println(name, age, isCool, name2, age2, isCool2,id,pi,appName)
	fmt.Println(sum)
	printSomething(sum)
	fmt.Println("hello world")
}
