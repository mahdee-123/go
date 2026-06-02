package main

import "fmt"

// 1–5 পর্যন্ত সংখ্যার square channel এর মাধ্যমে পাঠাও।
func main() {
	ch := make(chan int)


	go func ()  {
		for i := 1; i <= 5; i++ {
		/// create goroutines 
		fmt.Println("sending..")
		go func () {
			ch <- i * i
		}()

	}
		close(ch)
	}() 
	
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("recieving... ")
	// 	fmt.Println(<-ch)
	// } 


	for v := range ch {
		fmt.Println("recieving...")
		fmt.Println(v)
	}

}