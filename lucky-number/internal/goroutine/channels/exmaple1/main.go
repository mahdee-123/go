package main

import (
	"fmt"
	"time"
)

func main() {
	// crate chanals
	ch := make(chan int)
	go func() {
		fmt.Println("child goroutine running...")
		time.Sleep(time.Second*8)
		ch <- 10
	}()
	go func () {
		fmt.Println("second goroutine is running....")
	}()
	fmt.Println("main goroutine running...")
	
	fmt.Println("main is working again....")
	value := <- ch
	fmt.Println(value)
	
}