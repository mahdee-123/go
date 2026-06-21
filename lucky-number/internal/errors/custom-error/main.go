package main

import (
	"fmt"
	"log"
)

type MyError struct {
	code int
	msg  string
}

func (e MyError) Error() string {
	return fmt.Sprintf("[Code %d]", e.code)
}

func main() {
	err := MyError{code: 404, msg: "Not Found"}
	log.Fatal(err)
	fmt.Println("we go some isseue" + err.Error())
	fmt.Println(err)
}