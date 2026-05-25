package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

// Custom Errors 🔥
type ValidationsError struct {
	Field string
}

func (v ValidationsError) Error() string {
	return fmt.Sprintf("field %s is invalid", v.Field)
}
func main() {


	 result , err:= divide(10,0)
	 if err != nil {
		fmt.Println(err)
	 } else {
		fmt.Println(result)
	 }

		emailError := ValidationsError {
			Field: "email",
		}

		fmt.Println(emailError)

}