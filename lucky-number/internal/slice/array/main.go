package main

import "fmt"

// When our clients don't respond to a message, they can be reminded with up to 2 additional messages.

// Complete the getMessageWithRetries function. It takes three strings and returns:

// An array of 3 strings
// An array of 3 integers
// The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.

// The integers in the integer array represent the cost of sending each message. The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

// "hello" costs 5
// "world" costs 5, adding "hello" makes total cost 10 (5 + 5)
// "!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)
func getMessageWithRetries(primary, secondary, tertiary string) (messages [3]string, costs [3]int) {

	messages = [3]string{primary, secondary, tertiary} 

	costs[0] = len(primary)
	costs[1] = costs[0] + len(secondary) 
	costs[2] = costs[1] + len(tertiary)

	return	
}

func main() {
    // var numbers [5]int = [5]int{1, 2, 3, 4, 5}
    // fmt.Println(numbers)

    var slice []int 

    if slice == nil {
        fmt.Println("the zeroth value of slice is nil")
    }
    fmt.Println(slice)
    // ⚠️ এখানে return value store করো নি
    getMessageWithRetries("hi whats up", "are you there", "??")
    
    // ✅ এভাবে করলে ভালো:
    msgs, costs := getMessageWithRetries("hi whats up", "are you there", "??")
    fmt.Println("Messages:", msgs)
    fmt.Println("Costs:", costs)


    // underlaying array 
    numbers := []int{1,2,3,4,5}
    part := numbers[1:3]
    part[0] = 99
    fmt.Println(numbers)
    fmt.Println(part[0])
    fmt.Println(part[1])
}