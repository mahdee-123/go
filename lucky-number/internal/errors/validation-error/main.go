package main

import "fmt"

type ValidationError struct {
	Field   string
	message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("validation failed for %s %s", v.Field, v.message)
}
func validateUser(name string, age int) error {
	if name == "" {
		return ValidationError{Field: "Name", message: "Name is empty"}
	}
	if age <= 0 {
		return ValidationError{Field: "Age", message: "Age is invalid"}
	}
	return nil
}

func main() {


	// Test 1: Empty name
	err := validateUser("", 25)
	if err != nil {
		fmt.Println("Error:", err)
		// Error: validation failed for 'name': name cannot be empty
	}

	// Test 2: Age too low
	err = validateUser("রহিম", 15)
	if err != nil {
		fmt.Println("Error:", err)
		// Error: validation failed for 'age': age must be at least 18
	}

	// Test 3: Valid
	err = validateUser("আলী", 25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Validation passed!")
	}
}