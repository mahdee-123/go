package main

import (
	"fmt"
	"unicode"
)

// Password Strength
// As part of improving security, Textio wants to enforce a new password policy. A valid password must meet the following criteria:

// At least 5 characters long but no more than 12 characters.
// Contains at least one uppercase letter.
// Contains at least one digit.
// A string is really just a read-only slice of bytes. This means that you can use the same techniques you learned in previous lessons to iterate over the characters of a string.

// Assignment
// Implement the isValidPassword function by looping through each character in the password string. Make sure the password is long enough and includes at least one uppercase letter and one digit.

// Assume passwords consist of ASCII characters only.

// func isValidPassword(password string) bool {
// 	characterCount, digitCount , upperCaseLatercount := 0, 0, 0

// 	for _, char := range password {
// 		characterCount ++
// 		if char >= 'A' && char <= 'Z' {
// 			upperCaseLatercount++
// 		}
// 		if char >= '0' && char <= '9' {
// 			digitCount++
// 		}
// 	}

// 	return characterCount >= 5 && characterCount <= 12 && upperCaseLatercount >= 1 && digitCount >= 1

// }

/// solvig with built in fucntions

func isValidPassword(password string) bool {
	
	/// 
	if len(password)  < 5 || len(password) > 12  {
		return false
	}

	hasDigit := false
	hasUpper := false 

	for _, char := range password {
		if unicode.IsDigit(char) {
			hasDigit = true 
		}
		if unicode.IsUpper(char) {
			hasUpper = true 
		}

		if hasDigit && hasUpper {
			return true
		}

		
	}

	return  hasDigit && hasUpper 
}

func main() {

	fmt.Println(isValidPassword("password"))
	fmt.Println(isValidPassword("password1A"))

}