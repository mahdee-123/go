package main

import (
	"fmt"
	"strings"
)

// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// // func canSendMessage(mToSend messageToSend) bool {
// // 	if mToSend.sender.name == "" ||
// // 		mToSend.recipient.name == "" ||
// // 		mToSend.sender.number == 0 || mToSend.recipient.number == 0 {
// // 		return false
// // 	}

// // 	return true
// // }

// // te
// func canSendMessage(mToSend messageToSend) bool {
// 	return mToSend.sender.name != "" &&
// 	       mToSend.sender.number != 0 &&
// 	       mToSend.recipient.name != "" &&
// 	       mToSend.recipient.number != 0
// }
// func main() {

// 	user1 := user{"", 1234}
// 	user2 := user{"sadia bibi", 3453}

// 	messageToSend1 := messageToSend{"hello", user1, user2}
// 	messageToSend2 := messageToSend{"hello", user2, user1}

// 	fmt.Println(canSendMessage(messageToSend1))
// 	fmt.Println(canSendMessage(messageToSend2))
// }

// func main() {
// 	mycar := struct {
// 		brand string
// 		model string
// 	} {
// 		brand: "toyota",
// 		model: "v4",
// 	}

// 	fmt.Println(mycar)
// }

// type Person struct {
// 	name string
// 	number int
// 	educationalStatus educationalStatus
// }

// type educationalStatus struct {
// 	hsc bool
// 	ssc bool
// }

// func main() {
// 	p := Person{name: "Arnob", number: 1234, educationalStatus: educationalStatus{hsc: true, ssc: false}}

// 	fmt.Println(p.hsc)
// }

// type animal struct {
// 	name string
// }

// type Animal interface {
// 	Speak()
// }

// func (a animal) Speak() {
// 	fmt.Println("animal sound...")
// }

// type Dog struct {
// 	animal
// }

// func (a Dog) Speak() {
// 	fmt.Println("wooof...")
// }

// type Cat struct {
// 	animal
// }

// func (a Cat) Speak() {
// 	fmt.Println("meow...")
// }

// func makeNoisse(a Animal) {
// 	a.Speak()
// }

// func main() {
// 	d := Dog{ animal: animal{name: "animal"}}
// 	c := Cat{ animal: animal{name: "animal"}}
// 	makeNoisse(d)
// 	makeNoisse(c)
// }

// type authenticationInfo struct {
// 	username string
// 	password string
// }

// // create the method below
// func (a authenticationInfo) getBasicAuth() string {
// 	return fmt.Sprintf("Authorization: Basic %s:%s",a.username, a.password)
// }

// func main() {
// 		u1 := authenticationInfo{username: "arnob", password: "1234"}
// 		fmt.Println(u1.getBasicAuth())
// }

// type Membership struct {
// 	Type string
// 	MessageCharLimit int
// }

// type User struct {
// 	Name string
// 	Membership
// }

// func newUser(name string, membershipType string) User {

// 	var MessageCharLimit int
// 	if membershipType == "premium" {
// 		MessageCharLimit = 1000
// 	} else {
// 		MessageCharLimit = 100
// 	}

// 	return User{
// 		Name: "mahdee",
// 		Membership: Membership{
// 			Type: membershipType,
// 			MessageCharLimit: MessageCharLimit ,
// 		},
// 	}
// }

// Assignment
// Create a SendMessage method for the User struct.

// It should take a message string and messageLength int as inputs.

// If the messageLength is within the user's MessageCharLimit, return the original message and true (indicating the message can be sent), otherwise, return an empty string and false.







func (user User) SendMessage(message string , messageLength int) (msg string, ok bool) {
	if messageLength <= user.MessageCharLimit {
		return message, true
	}
	
	return "", false
}
type User struct {
	Name string
	Membership 
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

// func newUser(name string, membershipType string) User {
// 	membership := Membership{Type: membershipType}
// 	if membershipType == "premium" {
// 		membership.MessageCharLimit = 1000
// 	} else {
// 		membership.Type = "standard"
// 		membership.MessageCharLimit = 100
// 	}
// 	return User{Name: name, Membership: membership}
// }
type Mytype int

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}



func removeProfanity(message *string) {

	var updatedString string 

	updatedString = strings.ReplaceAll(*message, "fubb", "***")
	updatedString = strings.ReplaceAll(*message, "shiz", "***")
	updatedString = strings.ReplaceAll(*message, "witch", "***")

	*message = updatedString
}


type Person struct {
	Name string 
	Age  int
	shirtSize string
}

func (p *Person) updateSize(size string) {
	p.shirtSize = size 
}


func printPerson(p *Person) {
	fmt.Println(p.Name)
}


type Analytics struct {
    MessagesTotal int
}
func main() {
	
	p1 := Person{
		Name: "arnob",
		Age: 10,

	}

	p1.updateSize("xl")
	fmt.Println(p1)

	analytics := &Analytics{
		MessagesTotal: 100,
	}



	msgTotal := (*analytics).MessagesTotal
	fmt.Println(msgTotal)
	x := 10 
	p := &x 

	fmt.Println(*p)
	*p = 20
	fmt.Println(x) 

	var a Mytype = 10
	var y *Mytype = &a
 
	fmt.Printf("%T", y)

	value := getMessageText(Message{Recipient: "arnob", Text: "hello"})
	fmt.Println(value)


	var emptyPointer *int 

	fmt.Println(emptyPointer)
	message := "hello fubb shiz"
	removeProfanity(&message)
	fmt.Println(message)

	person1 := Person{
		Name: "arnob",
		Age: 10,
	}
	printPerson(&person1)




}