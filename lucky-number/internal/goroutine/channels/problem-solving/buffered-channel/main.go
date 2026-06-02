// Problem:
// Capacity 3 এর buffered channel তৈরি করে 3টি value পাঠাও।
package main 

import "fmt"

func main() {
	ch := make(chan int,3)
	ch <- 10
	ch <- 20
	ch <- 30 
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}