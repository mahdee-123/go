package main

import (
	"fmt"
)
type Person struct {
	Name string
	Age int
}
func main() {
		// empty interface is means no fixed data types,, that can be anything (like ts any)
		var i interface {}
		i = 10
		fmt.Println(i)
		i = true
		fmt.Println(i)
		i = "hello world"
		fmt.Println(i)

		var p interface {}


		p = Person{Name: "Arnob", Age: 10}
		

		// but there comes a issue, when youu want to use.. data type specific methods
		// fmt.Println(len(i)) // it will give you error
		// for this you have to do type asertion

		// type assertion
		s := i.(string)
		fmt.Println(len(s))

		// but In newer Go versions, people often use: any

		type any = interface {}
		var value any 
		value = 10
		fmt.Println(value)
		value = true
		fmt.Println(value)
		value = "hello world"
		// type assertion 
		v, ok := value.(string)
		// fmt.Println(len(v))
		
		if ok == false {
			fmt.Println("type assertion is wrong")
		} else {
			fmt.Println(v)
		}

		// / type assertion in struct 
		person := p.(Person) 

		fmt.Println(person.Name)
}
