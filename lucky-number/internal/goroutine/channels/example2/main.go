package main

import "fmt"

func worker(ch chan string) {
	ch <- "i love you"
}


func main() {

	ch := make(chan string)
 	go worker(ch)
	msg := <- ch
	fmt.Println(msg)
}