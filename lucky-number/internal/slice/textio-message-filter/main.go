package main

import "fmt"

// Message Filter
// Textio is introducing a feature that allows users to filter their messages based on specific criteria. For this feature, messages are categorized into three types: TextMessage, MediaMessage, and LinkMessage. Users can filter their messages to view only the types they are interested in.

// Assignment
// Your task is to implement a function that filters a slice of messages based on the message type.

// Complete the filterMessages function. It should take a slice of Message interfaces and a string indicating the desired type ("text", "media", or "link"). It should return a new slice of Message interfaces containing only messages of the specified type.

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}


func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}


func filterMessages(messages []Message, filterType string) []Message {
	filteredMessages := []Message{}
	for _, message := range messages {
		if message.Type() == filterType {
			filteredMessages = append(filteredMessages, message)
		} 
	}
	return filteredMessages
}
func main()  {

	messages := []Message{
		TextMessage{Sender: "Alice", Content: "Hello, world!"},
		MediaMessage{Sender: "Bob", MediaType: "image", Content: "https://example.com/image.jpg"},
		LinkMessage{Sender: "Charlie", URL: "https://example.com", Content: "Check out this link!"},
		TextMessage{Sender: "Alice", Content: "Goodbye, world!"},
		MediaMessage{Sender: "Bob", MediaType: "video", Content: "https://example.com/video.mp4"},
		LinkMessage{Sender: "Charlie", URL: "https://example.com", Content: "Check out this video!"},
		TextMessage{Sender: "Alice", Content: "Hello, world!"},
	}

	result := filterMessages(messages	, "text")
	fmt.Println(result)

}