package main

import "fmt"

func add(a, b int, ch chan int) {
	ch <- a + b
}
func main() {
	ch := make(chan int)
	go add(10, 20, ch)
	fmt.Println(<-ch)
}