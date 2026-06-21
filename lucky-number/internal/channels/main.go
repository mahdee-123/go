package main

import (
	"fmt"
	"time"
)

// Key Properties of channels
// - 	Type-safe			শুধু একটা type এর data যাবে ||	Water pipe তে শুধু water, oil pipe তে শুধু oil
// -	Blocking	   		Sender wait করে receiver না আসা পর্যন্ত || 	হাতে হাত দিয়ে deliver - দুইজন লাগবে
// -	Thread-safe			Multiple goroutines safely use করতে পারে	 || Built-in lock আছে
// -	FIFO					First In First Out || 	Queue - আগে ঢুকলে আগে বের হবে
// - 	var a string = "hello world!"
// var data string


// ekta way hocce earn money .... good proposal 
func main() {


	// ch := make(chan int)
	// names := make(chan string)
	// data := make(chan  bool)

	// type Person struct {
	// 	Name 	string 
	// 	Age 	int
	// }


	// persons := make(chan Person) 
	// var value int

	// ch1 := make(chan int)
	
	// go func ()  {
	// 	 value = <-ch1   
	// }() 
	     
	// ch1 <- 42  


	// fmt.Println(value)

	// ch := make(chan int)
	// go func()  {
	// 	ch <- 10 
	// }()

	// fmt.Println(<-ch)

		
	// ch := make(chan int,3)

	// ch <- 10
	// ch <- 11

	// close(ch)

	// for {
	// 	value, ok := <- ch
	// 	if ok {
	// 		fmt.Println(value)
	// 	}else {
	// 		break
	// 	}
	// }
	// for v:= range ch {
	// 	// what it does internally 
	// 	// - recieve data....continiously ...
	// 	// - if ok become false then break the loop autometically ...
	// 	fmt.Println(v)
	// }

	// fmt.Println("loop breaked")

	ch1 := make(chan int,1)
	ch2 := make(chan int,1)


	go func ()  {
		time.Sleep(2 * time.Second)
		ch1 <- 10
	}()

	go func ()  {
		time.Sleep(2 * time.Second)
		ch2 <- 10
	}()

	select {
	case v:= <- ch1: 
		fmt.Println(v)
	case v:= <- ch2: 
		fmt.Println(v)
	case <-time.After(1 * time.Second):
		fmt.Println("tiemout")
	}

	// দুইটা goroutine একে অপরকে message পাঠাবে
	// Output: Ping! Pong! Ping! Pong! ... (10 times)


	// ping , pong , ping , ping 
	pong := make(chan string)
	ping := make(chan string)

	go func ()  {

		// this is break... up 
		for range 10 {
			ping <- "Ping!"
		}

		close(ping)
	}()

	for v:= range ping {
		pong <- v
		fmt.Println()
	}


}


