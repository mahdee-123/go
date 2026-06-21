package main

import "fmt"

func main() {
	m1 := make(map[string]int)


	// Insert an Element
	m1["key1"] = 10
	m1["key2"] = 20

	fmt.Println(m1)

	// get an Element

	getAnElement := m1["key1"]
	fmt.Println(getAnElement) 

	// Delete an Element 
	fmt.Println(m1)
	delete(m1,"key1")
	fmt.Println(m1)

	// Check If a Key Exists or Not 
	elem, ok := m1["key1"]
	if ok == false {
		fmt.Println("key not found") 
	} else {
		fmt.Println(elem)
	}
}