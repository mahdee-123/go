package main

import (
	"fmt"
	"strings"
)


// default types 

// Message Tagger
// Textio needs a way to tag messages based on specific criteria.

// Assignment
// Complete the tagMessages function. It should take a slice of sms messages, and a function (that takes an sms as input and returns a slice of strings) as inputs. And it should return a slice of sms messages.
// It should loop through each message and set the tags to the result of the passed in function.
// Be sure to modify the messages of the original slice using bracket notation messages[i].
// See the tip below on how the strings package could be used here.

// Complete the tagger function. It should take an sms message and return a slice of strings.
// Return an initialized slice, even if no tags match. No nil slices.
// For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
// For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
// Regardless of casing just means that even "uRGent" or "SALE" should trigger the tag.

// Example usage:

// messages := []sms{
// 	{id: "001", content: "Urgent! Last chance to see!"},
// 	{id: "002", content: "Big sale on all items!"},
// 	// Additional messages...
// }
// taggedMessages := tagMessages(messages, tagger)
// // `taggedMessages` will now have tags based on the content.
// // 001 = [Urgent]
// // 002 = [Promo]

// Tip
// The go strings package, specifically the Contains and ToLower functions, can be very helpful here!




type sms struct {
	id      string
	content string
	tags    []string
}



func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for index := range messages {
		messages[index].tags = tagger(messages[index])
	}
	return messages
}

func tagger(msg sms) []string {

	tags := []string{}
	
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
	
}

func main() {

	messages := []sms{
		{id: "001", content: "Urgent! Last chance to see!"},
		{id: "002", content: "Big sale on all items!"},
		{id: "003", content: "Hello, world!"},
		{id: "004", content: "Goodbye, world!"},
	}
	
	taggedMessages := tagMessages(messages, tagger)
	// fmt.Print(messages)
	fmt.Print(taggedMessages)
	// `taggedMessages` will now have tags based on the content.
	// 001 = [Urgent]
	// 002 = [Promo] 
}
