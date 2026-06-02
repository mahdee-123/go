package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func ()  {
		ch <- 10
		fmt.Println("sent")
	}()
	

	time.Sleep(time.Second*2)
	msg := <- ch
	fmt.Println(msg)
	
	fmt.Println("main goroutine")
	time.Sleep(time.Second*2)
	
}						