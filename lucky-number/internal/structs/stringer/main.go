package main

import "fmt"


// type Stringer interface {
// 	String() string
// }

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	l := Person{Id: 1, Name: "Arnob", Age: 20}
	fmt.Println(l)
}



func (p Person) String() string {
	return fmt.Sprintf("Person Type\nID : %d\nName : %s\nAge :%d",p.Id, p.Name, p.Age)
}