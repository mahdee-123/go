// Textio now has a system to process different types of notifications: direct messages, group messages, and system alerts. Each notification type has a unique way of calculating its importance score based on user interaction and content.
package main

import "fmt"


type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}


// all methods 
func (dm directMessage) importance() int {
	if dm.isUrgent == true {
		return 50
	} else {
		return dm.priorityLevel
	}
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}
func processNotification(n notification) (string, int) {
	switch n.(type) {
	case directMessage:
		return n.(directMessage).senderUsername, n.(directMessage).importance()
	case groupMessage:
		return n.(groupMessage).groupName, n.(groupMessage).importance()
	case systemAlert:
		return n.(systemAlert).alertCode, n.(systemAlert).importance()
	default:
		return "", 0 
	}
}

func main() {

	notifications := []notification {
		directMessage{
			senderUsername: "user1",
			messageContent: "Hello, world!",
			priorityLevel:  10,
			isUrgent:       true,
		},
		groupMessage{
			groupName:      "Group A",
			messageContent: "Group message",
			priorityLevel:  20,
		},
		systemAlert{
			alertCode:      "ALERT-001",
			messageContent: "System alert",	
	} } 

	for _, notification := range notifications {
		 fmt.Println(processNotification(notification))
	}
}