package main

import "fmt"

func main() {

	hits := map[string]map[string]int{
		"Bangladesh": {
			"/home":    100,
			"/about":   50,
			"/contact": 30,
		},
		"India": {
			"/home":  200,
			"/about": 80,
		},
	}

	hits["Pakistan"] = map[string]int{
		"/home":    100,
		"/about":   50,
		"/contact": 30,
	}

	// keys in golang

	m1 := make(map[string]int)

	m1["key1"]++ 
	fmt.Println(m1)
}
