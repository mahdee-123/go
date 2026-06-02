package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {

	for num := range numChan {
		fmt.Println("processing ", num)
	}
}
func main() {

	numChan := make(chan int)
	go processNum(numChan)

	for  {
		numChan <- rand.Intn(100)
		time.Sleep(time.Second * 1)
	}

}