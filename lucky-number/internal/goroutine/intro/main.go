package main

import (
	"fmt"

	
	"time"
)

func task() {
	for i := 1; i <= 5; i++ {
		fmt.Println("task: ", i)
		time.Sleep(time.Millisecond*500)
	}
}

func main() {

	go task()

	for i := 1; i <= 5; i++ {
		fmt.Println("main: ", i)
		time.Sleep(time.Millisecond*500)
	}

}