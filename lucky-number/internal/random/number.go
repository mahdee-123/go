package random

import (
	"fmt"
	"math/rand"
)

type Celcius float32

func (c Celcius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}

const  (
	Red = iota + 4
	Blue 
	Yellow 
)

func Number() int {


	var numbers = make([]int, 2, 19)
	fmt.Println(cap(numbers))
	numbers = append(numbers, 4)
	fmt.Println(numbers)

	tem := Celcius(50) 

	if num := rand.Intn(100); num > 50 {
		fmt.Println("Greater than 50")
	}
	i := 0
	
	for  i<10 {
		fmt.Println(i)
		i++
	} 

	// infinite loop 

	nums := []int{10, 20, 30}


	for index, value := range nums {
		fmt.Println(index, value)
	}


	day := 2
	switch day {
	case 1 : 
		fmt.Println("Monday")
	case 2 : 
		fmt.Println("Tuesday")
	case 3 : 
		fmt.Println("Wednesday")
	default : 
		fmt.Println("Unknown")
	}
	
	fmt.Println(tem.ToFahrenheit())


	fmt.Println(Red,Blue, Yellow)

	rolls := []int{1,2,3,4,5,6}
	rolls = append(rolls, 4)
	fmt.Println(rolls)
	marks := [...]int{3,3,31,3}
	fmt.Println(marks)
	return rand.Intn(100)
}

