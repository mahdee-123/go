package main

import (
	"fmt"
	"sync"
)

var count int = 0


var mu sync.Mutex 

func main() {

	var wg sync.WaitGroup
	for range 1000 {

		wg.Add(1)
		// race conditions
		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Println("hello")
			mu.Unlock()
		}()

	}

	wg.Wait()

	fmt.Println(count)
}