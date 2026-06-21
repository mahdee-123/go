package main

import (
	"fmt"
)

// We send a lot of text messages at Textio, and our API is getting slow and unresponsive.

// If we know the rough size of a slice before we fill it up, we can make our program faster by creating the slice with that size ahead of time so that the Go runtime doesn't need to continuously allocate new underlying arrays of larger and larger sizes. By setting the length, the slice can still be resized later, but it means we can avoid all the expensive resizing since we know what we'll need.

// Complete the getMessageCosts() function. It takes a slice of messages and returns a slice of message costs.

// Preallocate a slice for the message costs of the same length as the messages slice.
// Fill the costs slice with costs for each message. The cost in the cost slice should correspond to the message in the messages slice at the same index. The cost of a message is the length of the message multiplied by 0.01.

func getMessageCosts(messages []string) []float64 {
	arrayLength := len(messages)
	costs := make([]float64,arrayLength) // preallocate a slice for the message costs of the same length as the messages slice.

	for index, message := range messages {
		costs[index] = float64(len(message)) * 0.01
	}

	return costs
}

/// Assignment
// We need to sum up the costs of all individual messages so we can send an end-of-month bill to our customers.

// Complete the sum function to return the sum of all inputs.

// Take note of how the variadic inputs and the spread operator are used in the test suite.

func sum(nums ...int) (sum int) {
	
	for _, num := range nums {
		sum += num
	}

	return 
}


/// Assignment
// We've been asked to add a feature that extracts costs for a given day.

// Complete the getDayCosts() function using the append() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day's costs. A day may have multiple costs.

// If there are no costs for that day, return an empty, non-nil slice.
type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	
	var allCosts []float64
	for _, cost := range costs {
		if cost.day == day {
			allCosts = append(allCosts, cost.value)
		}
	}

	return allCosts
}

func main() {

	messeages := []string{"hi" , "hello", "ami "}
	getMessageCosts(messeages)
	fmt.Println(sum(12,3,4))

	// 
	getDayCosts([]cost{}, 1)


}