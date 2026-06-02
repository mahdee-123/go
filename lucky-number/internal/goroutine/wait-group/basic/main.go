package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println("goroutine start executing..")
	time.Sleep(1 * time.Second)
	fmt.Println("finish exucuting myFunc") 
	
}
func main() {
	fmt.Println("execution start here....")
	var wg sync.WaitGroup
	wg.Add(1)
	go myFunc(&wg)
	wg.Wait() // block until zero
	fmt.Println("finish the the programmed here..")

}