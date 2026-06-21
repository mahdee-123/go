package main

import "errors"

// # divide(a, b int) (int, error) — b==0 হলে error, properly check করে call করো।


func divide(a,b int) (result int, err error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	result = a / b
	return 
}
func main() {

	divide(1,0)
	divide(2,1)

}