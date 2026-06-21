// unbuffered channel
// কখন: সবসময় শুরুতে
// 1. একটা unbuffered channel বানাও (chan string)। একটা goroutine থেকে একটা message পাঠাও, main এ receive করো। লক্ষ্য করো — receive না হওয়া পর্যন্ত send block থাকে।

package main

// package main -->> ...arg

import "fmt"


func main() {

	var ch1 = make(chan string)

	go func() {
		ch1 <- "hello"
	}()

	value , ok := <- ch1

	if  !ok {
		fmt.Println("channel is closed")
	}
	
	fmt.Println(value)


	fmt.Println("main executes")
	
}