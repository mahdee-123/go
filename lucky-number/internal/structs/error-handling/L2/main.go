package main

import (
	"errors"
	"fmt"
)

func main() {
	msg, err := registerUser("arham", -19)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}
}

func registerUser(name string, age int) (string, error) {
	if name == "" {
		return "", errors.New("name is empty")
	} 
	if age <= 0 {
		return "", errors.New("age cannot be lower or equal to 0")
	}
	return "user registered successfully", nil
}
