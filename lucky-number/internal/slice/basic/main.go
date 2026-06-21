package main

import "fmt"

type Slice struct {
	pointer  *int
	lenght   int
	capacity int
}

func main() {

	var slice1 []int
	slice2 := []int{}

	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)


	original := []int{1,2,3,4}
	part := original[0:2:2]

	part = append(part, 999)
	fmt.Println(original)

	fmt.Println(cap(original))

	// using make 
	// Form 2: Length + Capacity
	s1 := make([]int,3,4)
	fmt.Println(s1)

	// Form 1: Length only

	s2 := make([]string, 3) // len = 3, cap = 3
	fmt.Println(s2)
	// Form 3: Empty slice (length = 0)

	s3 := make([]int, 0)
	fmt.Println(s3)

	// slice literal 
	mySlice := []string{"I", "love", "go"} 
	fmt.Println(mySlice) 



	// 

	a := make([]int,3)
	fmt.Println(len(a))
	
	fmt.Println(cap(a))
	fmt.Println(a)

	b := append(a,4)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("address of b:", &b[0])


	c := append(a,5)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)


	fmt.Println("starting....")
	i := make([]int, 3,8) 
	// underlying array1 : [0 0 0 _ _ _ _ _]
	// i : len = 3, cap = 8 
	// fmt.Println(len(i))
	// fmt.Println(cap(i))


	j := append(i,4)
	// underlying array1 : [0 0 0 4 _ _ _ _]
	// j : len = 4, cap = 8 
	fmt.Println(j)
	fmt.Println(&j[0])

	g := append(i,5)
	// underlying array1 : [0 0 0 5 _ _ _ _]
	// j : len = 4, cap = 8 
	fmt.Println(g)
	fmt.Println(&g[0])

	fmt.Println("i:", i)
	fmt.Println("j:", j)
	fmt.Println("g:", g)

}