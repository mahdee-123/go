package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	// Full-time employee
	ft := fullTime{
		name:   "Alice",
		salary: 50000,
	}

	// Contractor
	ct := contractor{
		name:         "Bob",
		hourlyPay:    50,
		hoursPerYear: 2000,
	}

	// উভয়কেই employee হিসেবে ব্যবহার করা যাচ্ছে ✅
	employees := []employee{ft, ct}

	for _, emp := range employees {
		fmt.Printf("Name: %s, Salary: %d\n",
			emp.getName(),
			emp.getSalary())
	}
}


