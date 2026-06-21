package main

import (
	"fmt"
	"strconv"
)

// atoi == ascii to integer


func main() {
	userInput := "880139304034"

	num, err  := strconv.Atoi(userInput)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T", num )
		fmt.Println()
		fmt.Println(num)
	}

}