package main

import (
	"fmt"
	"time"
)

func worker(stop *chan bool)  {
	for {
		select {
		case <- *stop: 
			fmt.Println("stoppped this goroutine..")
			return 

		default : 
			fmt.Println("working..")
			time.Sleep(time.Second)
		} 
	}
}
func main() {
	stop := make(chan bool)
	stop <- true
	go worker(&stop)
	time.Sleep(time.Second * 5)
	fmt.Println("main finished working....")	
}