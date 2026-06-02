package main

import (
	"fmt"
	"sync"
	"time"
)

type Product struct {
	ID           int
	Availability string
	Price        float64
}

func main() {

	var wg sync.WaitGroup
	product1 := &Product{ID: 1}

	wg.Go(func() {
		fmt.Println("Fetching availability for product one......")
		time.Sleep(100 * time.Millisecond)
		product1.Availability = "In Stock"
		fmt.Println("Availability fetched!")
	}) 

	wg.Go(func() {
		fmt.Println("Fetching price for product one......")
		time.Sleep(100 * time.Millisecond)
		product1.Price = 100.0
		fmt.Println("Price fetched!")
	}) 
	
	fmt.Println("router is waiting for subgraph responsse....")
	wg.Wait()

	fmt.Printf("Successfully fetched data for Product ID %d: %+v\n", product1.ID, *product1)

}