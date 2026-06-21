package main

import "fmt"
// coding challenge..... 
func main() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	value ,ok:= timeZone["sdfs"]
	fmt.Println(value, ok)
}