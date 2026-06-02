package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["key1"] = 10
	m["key2"] = 20

	fmt.Println(m)


	v1 := m["key1"]
	fmt.Println(v1) 

	len := len(m)
	fmt.Println(len) 

	// delete(m,"key1")
	// fmt.Println(m)

	// clear(m)
	// fmt.Println(m)

	value, prs := m["key1"]
	fmt.Println(value, prs) 

	map1 := map[string]int{
		"key1": 10,
		"key2": 20, 
	}

	map2 := map[string]int{
			"key1": 10,
			"key2": 20, 
	}

	fmt.Println(map2)
	fmt.Println(map1 == nil) 

	if maps.Equal(map1, map2) {
		fmt.Println("same maps")
	}


	map3 := map[string]int{
		"key1": 10,
		"key2": 20, 
		"key3": 30, 
		"key4": 40, 
		"key5": 50, 
		"key6": 60, 
		"key7": 70, 
	}

	fmt.Println("iteratoin on map3")
	for _, v := range map3 {
		fmt.Println(v)
	}

	
}