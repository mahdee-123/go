package main

import (
	"errors"
	"fmt"
)

func divide(dividend, divisor float64) (result float64, err error) {
	if divisor == 0 {
		err = errors.New("cannot divide by zero ")
	} else {
		result = dividend / divisor
	}
	return

}
func main() {

	result, err := divide(1,0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}