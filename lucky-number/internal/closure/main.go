package main

import "fmt"

func counter() func() int {
	counter := 0

	return func() int {
		counter++
		return counter
	}
}

func main() {

	c1 := counter()
	c2 := counter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())

}