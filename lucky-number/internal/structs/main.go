package main

import "fmt"

// most powerful data types : structs
// not same as object or oop
// go is not a obejct oriented language
/// struct is data type that let us group related data together
// struct is custom data type
// struct can have differetnt data types together
// in most cases struct will represent business entities

// write a struct
type Person struct {
	Name string 
	Age int
	Email string
}


type Student struct {
	Name string
	Id int 
	Email string
}


func registerStudent(name string, id int, email string ) Student {
	return Student {
		Name : name, 
		Id : id, 
		Email : email,
	}
}	


func New(name string , age int, email string) *Person {
	return &Person {
		Name : name, 
		Age : age, 
		Email : email,
	}
}


func retrieveEmail(name string) string {
	return name + "@gmail.com"
}
// the variables inside struct are called fields
func main() {
	// Using a struct literal
	p1 := Person {
		Name : "Arnob", 
		Age : 10, 
		Email : "a@b.com",
	}
	// Using empty declaration and assignment
	var p2 Person
	p2.Name = "mahdee"
	p2.Age = 30
	p2.Email = retrieveEmail(p2.Name)
	fmt.Println(p2)
	fmt.Println(p1)

	// using factory function
	registerStudent("Arnob", 10, "a@b.com")

	p3 := New("Arnob", 35, "as@bsd.com")
	p4 := New("df", 2, "as@33.com")
	fmt.Println(p3)
	fmt.Println(p4)

	

	// Using the built-in new function
	p5 := new(Person)
	// p5.Name = "akkhor"
	// p5.Age = 39
	// p5.Email = "a@bsdf.com"
	fmt.Printf("%p",p5)
	
}