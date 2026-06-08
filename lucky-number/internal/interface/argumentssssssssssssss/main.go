package main

import "fmt"

func PrintAll(values ...interface{}) {
	for _, value := range values {
		fmt.Println(value)
	}
}
func main() {

	PrintAll(
		10,"string",
	)
}
