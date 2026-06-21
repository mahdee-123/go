// 1টা producer, 3টা workers, 1টা collector
// Producer: 1-9 numbers generate করবে
// Workers: প্রতিটা number কে square করবে
// Collector: সব results print করবে

package main

import (
	"fmt"
	"time"
)

func producer(jobs chan <- int) {

	for i := 1; i <= 9; i++ {
		jobs <- i
		fmt.Printf("producer sent %d: ", i)
	}

	close(jobs)
	fmt.Println("producer done , job chanell closed!")

}

func worker(id int, jobs chan int , result chan int) {

	for j := range jobs {
		
		fmt.Printf("Worker %d got job: %d\n", id, j)
       
		time.Sleep(500*time.Millisecond)

		fmt.Printf("worker %d processing job %d\n", id, j)
		result <- j * j
	}

	 fmt.Printf("  Worker %d done (no more jobs)\n", id)
}


func main() {

	var jobs = make(chan int, 9)
	var results = make(chan int, 9) 


	go producer(jobs)

	for i := 1; i <= 3; i++ {
		go worker(i,jobs, results)
	}
	
}