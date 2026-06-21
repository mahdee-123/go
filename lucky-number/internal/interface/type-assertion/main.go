package main

import (
	"fmt"
	"strings"
)

func main() {
	var value interface{} = "hello"

	msg := value.(string)

	fmt.Println(strings.Split(msg, ""))

	var v2 interface{} = 10 

	v2I := v2.(int)
	fmt.Println(v2I * 2)

}