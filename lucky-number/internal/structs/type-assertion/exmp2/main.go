package main

import (
	"fmt"
	"math"
)


type Shape interface {
	Area() float64
}
type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}


func (r Rect) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rect) diogonal() float64 {
	return math.Sqrt(r.width * r.width + r.height * r.height)
}
func main() {


	var s1 Shape 
	s1 = Rect{width: 10, height: 10}
	var s2 Shape
	s2 = Circle{radius: 10}

	fmt.Println(s1.Area())
	fmt.Println(s2.Area())
	fmt.Println(s1.(Rect).width)
	fmt.Println(s2.(Circle).radius)
	v , ok := s1.(Circle)
	fmt.Println(v, ok)
	fmt.Println(v.radius)
	

}