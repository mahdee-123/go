package main

import (
	"fmt"
)

type Address struct {
	City string
	State string
}

type Childs struct {
	childs []string
}
// 9. Slice Inside Struct
type Person struct {
	Name  string
	Age   int
	Email string
 	Address
	Childs
}


type Student struct {
	Name string
	Id int
	Email string
	Sec rune 
	Marks int
}


type Teacher struct {
	CoureName string 
	DepartmentName string
}


func createTeacher(course string , department string) Teacher {
	return Teacher{
		CoureName : course,
		DepartmentName : department,
	}

}

func emailGenerate(s *Student) string {
	s.Email = s.Name + "@gmail.com"
	return s.Email
}



func main() {
	person1 := Person{
		Name:  "Arnob",
		Age:   20,
		Email: "a@b.com",
		Address: Address{
			City: "Sylhet",
		},
		Childs: Childs{
			childs: []string{"a", "b", "c"},
		},
	}
	fmt.Println(person1)
	person2 := Person{ 
		Name:  "Arnob",
		Age:   20,
		Email: "a@b.com",
	}

	// 2. Access Struct Field
	fmt.Println(person1.Age)
	fmt.Println(person2)
	// 3. Update Struct Field
	person1.Age = 40 
	fmt.Println(person1.Age)

	person3 := Person {}
	person3.Age = 20
	// person3.Name = "sakhawat"
	person3.Email = "sakhawat@gmail.com"
	fmt.Println(person3)

	// 6. Anonymous Struct

	user := struct {
		Name string
		Age int
	 }{
		Name: "Arnob",
		Age: 20,
	 }

	fmt.Println(user)

	// 7. Struct Comparison
	// fmt.Println(person1 == person2)

	// 8. Nested Struct
	fmt.Println(person1.City)



	// 10. Struct Slice
	var allStudents = []Student{
		{Name: "Arnob", Id: 1, Sec: 'A'},
		{Name: "Rahim", Id: 2, Sec: 'A'},
		{Name: "Karim", Id: 3, Sec: 'B'},
		{Name: "Sakib", Id: 4, Sec: 'B'},
		{Name: "Nayeem", Id: 5, Sec: 'C'},
		{Name: "Hasan", Id: 6, Sec: 'A'},
		{Name: "Jahid", Id: 7, Sec: 'C'},
		{Name: "Tamim", Id: 8, Sec: 'B'},
		{Name: "Rafi", Id: 9, Sec: 'A'},
		{Name: "Nabil", Id: 10, Sec: 'C'},
	}
	// 12. Struct Function Parameter
	for i := range allStudents {
			emailGenerate(&allStudents[i])

	}
	
	fmt.Println(allStudents)

	// 11. Loop Through Struct Slice
	for _, student := range allStudents {
		fmt.Println(student.Name, student.Id, student.Sec)
	}

}