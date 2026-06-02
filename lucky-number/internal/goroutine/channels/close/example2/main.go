package main

import "fmt"

func main() {

	ch := make(chan int,3)

	ch <- 10
	ch <- 20
	ch <- 30
	// sent all data to chanals
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
