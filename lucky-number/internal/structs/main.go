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
	 Person
	Id int 

}


func registerStudent(name string, email string, age int, id int ) Student {
	return Student {
		Person : Person {
			Name : name, 
			Age : age, 
			Email : email,
		},

		Id : id,
	}
}	


func New(name string , age int, email string) *Person {
	return &Person {
		Name : name, 
		Age : age, 
		Email : email,
	}
}
//  value reciever vs pointer reciever 
// value reciever
func (p Person) greet()	{
	fmt.Println("ki obostha", p.Name)
}

func (p *Person) changeName(newName string) {
	p.Name = newName
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
	student1 := registerStudent("Arnob", "a@b.com",10,1110 )
	fmt.Println("student : ",student1.Name)
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
		fmt.Println()


	p2.greet()
	fmt.Println(p2.Name)
	p2.changeName("raven returns")
	fmt.Println(p2.Name)

	testOverriding()
}