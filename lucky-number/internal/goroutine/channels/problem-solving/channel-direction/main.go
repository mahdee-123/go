package main

import "fmt"

func sender(ch chan<- int) {
	ch <- 10
}

func receiver(ch <-chan int) {
	fmt.Println(<-ch)
}
func main() {
	ch := make(chan int)
	go sender(ch)
	receiver(ch)
}