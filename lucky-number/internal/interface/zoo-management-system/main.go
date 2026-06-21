package main 

type Animal interface {
	Feed()
	Sleep()
	Wash() 
}

type Dog struct { name string }
type Cat struct { name string }
type Cow struct { name string }

func (d Dog) Feed() {}
func (d Dog) Sleep() {} 
func (d Dog) Wash() {}


func (d Cat) Feed() {}
func (d Cat) Sleep() {} 
func (d Cat) Wash() {}


func main() {
	
}