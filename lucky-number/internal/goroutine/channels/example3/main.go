package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		fmt.Println("A")
		ch <- "hello" // pause 
		fmt.Println("B")
	}()

	fmt.Println("C")
	time.Sleep(time.Second*2)
	msg := <- ch

	fmt.Println(msg)

	fmt.Println("D")

	// C 
	// A
	// B
	// hello 
	// d

	
}