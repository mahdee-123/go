package main

import (
	"fmt"
)

// declaere functions


func greet(name string)  {
	fmt.Printf("hello %s how are you", name)
}


// Short Parameters
func init() {
	fmt.Println("SECOND INIT")
}
func init() {
	fmt.Println("FIRST INIT")
}



//  Multiple Return Values 


func divide(a, b int) (int , int) {
	q := a / b 
	r := a % b 
	return q, r
}

// named return value

func rectangle(a, b int ) (Area int) {
	Area = a * b 
	return 
}

// 5. Variadic Functions (...args)
func sum(nums ...int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	return total
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	greet("arnob")
	a,b := divide(10, 3)
	fmt.Println(a,b )

	rectangle(3,4)
	result := sum(4,3,34,1)
	fmt.Println(result)

	// . First-Class Functions
	avg := func (nums ...int) float32 {
		return float32(sum(nums...))/float32(len(nums))
	}


	fmt.Println(avg(1,2,3,4,5))
	fmt.Println(counter()())
	fmt.Println(counter()())
	fmt.Println(counter()())
	fmt.Println(counter()())
	fmt.Println(counter()())






	defer func() {
		msg := recover()

		fmt.Println("recovered", msg)

	}()

	panic("something bad happened")

}