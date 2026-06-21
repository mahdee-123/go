package main

import "fmt"

type Notifier interface {
    Send(message string)
}

type Email struct{}
func (e Email) Send(msg string) {
    fmt.Println("📧 Email:", msg)
}

type SMS struct{}
func (s SMS) Send(msg string) {
    fmt.Println("📱 SMS:", msg)
}

type PushNotification struct{}
func (p PushNotification) Send(msg string) {
    fmt.Println("🔔 Push:", msg)
}

// একটা function সব notification এর জন্য!
func SendNotification(n Notifier, msg string) {
    n.Send(msg)
}

// অনেক user কে একসাথে notify করা
func NotifyUsers(users []Notifier, msg string) {
    for _, user := range users {
        user.Send(msg)
    }
}

func main() {
    // বিভিন্ন উপায়ে notify
    SendNotification(Email{}, "নতুন অফার!")
    SendNotification(SMS{}, "OTP: 1234")
    
    // একসাথে সবাইকে
    users := []Notifier{
        Email{},
        SMS{},
        PushNotification{},
    }
    NotifyUsers(users, "সিস্টেম মেইন্টেনেন্স")
}