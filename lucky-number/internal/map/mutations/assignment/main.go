// It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

// Complete the deleteIfNecessary function. The user struct has a scheduledForDeletion field that determines if they are scheduled for deletion or not.

// If the user doesn't exist in the map, return the error not found.
// If they exist but aren't scheduled for deletion, return deleted as false with no errors.
// If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.

package main

import (
	"errors"
	"fmt"
)


func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	targetUser, ok := users[name] 

	// If the user doesn't exist in the map
	if ok == false {
		return false, errors.New("user not found")
	}

	// If they exist but aren't scheduled for deletion
	if targetUser.scheduledForDeletion == false {
		return false, nil
	}

	// If they exist and are scheduled for deletion
	delete(users, name)
	return true, nil

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	m1 := make(map[string]user)

	m1["user1"] = user{name: "user1", number: 1, scheduledForDeletion: false}
	m1["user2"] = user{name: "user2", number: 2, scheduledForDeletion: true}
	m1["user3"] = user{name: "user3", number: 3, scheduledForDeletion: false} 

	result ,err := deleteIfNecessary(m1, "user2")
	if err != nil {
		fmt.Println(err) 
	} else  {
		fmt.Println(result) 
	}
}	