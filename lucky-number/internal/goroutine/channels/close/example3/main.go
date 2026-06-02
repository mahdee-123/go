package main

import "fmt"

func main() {

	ch := make(chan int)


	go func (){
		fmt.Println("child......")
		ch <- 10
		fmt.Println("ch 1 sent......")
		ch <- 30
		fmt.Println("ch 2 sent......")
		ch <- 20
		fmt.Println("ch 3 sent......")
		// sent all data to chanals
		close(ch)
	}()

	
	fmt.Println("main.....")
	for value := range ch {
		fmt.Println("range over ch.... ")
		fmt.Println(value)
	}
}
