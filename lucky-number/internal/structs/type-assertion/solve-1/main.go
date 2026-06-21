package main

import "fmt"

func main() {
	// এই সব data আছে
	values := []interface{}{
		"রহিম",
		25,
		3.14,
		true,
		42,
		"ঢাকা",
	}

	//  TODO: প্রতিটা value এর type check করো
	// এবং safe way ব্যবহার করো (ok variable)

	for _, value := range values {
		// তোমার code এখানে লিখবে
		// hint: যদি string হয় তাহলে print করো "Name: xxx"
		if str, ok := value.(string); ok {
			fmt.Println("Name: ", str)
		}
		// যদি int হয় তাহলে print করো "Age: xxx"
		if age, ok := value.(int); ok {
			fmt.Println("Age: ", age)
		} 
		// যদি float64 হয় তাহলে print করো "Height: xxx"
		if height, ok := value.(float64); ok {
			fmt.Println("Height: ", height)
		}
		// যদি bool হয় তাহলে print করো "Active: xxx"
		if active, ok := value.(bool); ok {
			fmt.Println("Active: ", active)
		} 
	}
}