package main

import (
	"fmt"
	"errors"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		// Error একটা value হিসেবে return হচ্ছে
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)

	// Error check করছি normal if condition দিয়ে
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	
	fmt.Println("Result:", result)
}