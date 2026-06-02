package main

import "fmt"

func main() {

	ch := make(chan int)
	close(ch) // can send data to channals
	ch <- 10
	for {
		value := <-ch
		fmt.Println(value)
	} // deadlock


}