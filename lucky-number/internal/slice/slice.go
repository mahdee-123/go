package main

import (
	"fmt"
	"reflect"
	"slices"
	"sort"
)

func reverse(nums []int)  {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func contains(nums []int , target int)  {
	for _, num := range nums {
		if num == target {
			fmt.Println("Found")
			return
		}
	}
}

// func removeDuplicate(nums []int) {
// 	unique := []int{}
// 	for _, num := range nums {
// 		fmt.Print(num)
// 	}
// 	fmt.Println(unique)
// }
// type slice struct {
// 	ptr *int 
// 	let int 
// 	cap int
// }



func main() {
	// Problem 1: একটি integer slice তৈরি করো এবং print করো
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	
	// Problem 2: Slice এর length বের করো
	fmt.Println(len(slice1))

	// Problem 3: Slice এর capacity বের করো
	fmt.Println(cap(slice1))
	
	
	// Problem 4: make() ব্যবহার করে slice তৈরি করো
	
	var slice2 = make([]int, 5, 10)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	fmt.Println(slice2)


	// Problem 5: Slice এর element access করো
	fmt.Println(slice1[0])

	// Problem 6: Slice এর element update করো
	slice1[0] = 10
	fmt.Println(slice1)


	// Problem 7: Slice iterate করো using range
	for index, value := range slice1 {
		fmt.Println(index,":", value)
	}


	// Problem 8: Slice এ নতুন element append করো

	slice1 = append(slice1, 6)
	fmt.Println(slice1)

	// Problem 9: একাধিক element append করো
	slice1 = append(slice1, 7,8,9,10)
	fmt.Println(slice1)

	// Problem 10: একটি slice আরেকটি slice এ append করো
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	// Problem 11: Slice থেকে sub-slice তৈরি করো
	subslice := slice1[1:3]
	fmt.Println(subslice)


	// Problem 12: First 3 elements বের করো
	firstThreeElement := slice1[:3]	
	fmt.Println(firstThreeElement)

	// Problem 13: Last 2 elements বের করো
	lastTwoELement := slice1[len(slice1)-2:]
	fmt.Println(lastTwoELement)

	// Problem 14: Slice copy করো
	slice3 := []int{4,3,5}
	slice4 := make([]int, len(slice3))
	copy(slice4, slice3)
	fmt.Println(slice3,slice4)


	// Problem 15: Slice reverse করো
	reverse(slice1)
	fmt.Println(slice1)


	// Problem 16: Slice clone করো
	original := []int{1,2,3,4,5,6,7,8,9,10}
	cloned := append([]int{}, original...)
	fmt.Println(cloned)
	
	// Problem 17: Slice compare করো 
	slice5 := []int{1,2,3,4,5,6,7,8,9,10}
	slice6 := []int{1,2,3,4,5,6,7,8,9,10}

	fmt.Println(reflect.DeepEqual(slice5,slice6))


	// Problem 18: Slice এর সব element এর sum বের করো
	var sum int  = 0;
	for _, value := range slice1 {
		sum += value
	}
	fmt.Println(sum)

	// Problem 19: Maximum value বের করো
	slice7 := []int{10, 50, 20, 80, 5}
	var max int = slice7[0]

	for _, value := range slice7 {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)


	// Problem 20: Minimum value বের করো
	var min int = slice7[0]

	for _, value := range slice7 {
		if value < min {
			min = value 
		}
	}
	fmt.Println(min)


	// Problem 21: Even numbers filter করো
	slice8 := []int{1, 2, 3, 4, 5, 6}
	evens := []int{}

	for _, value := range slice8 {
		if value%2 == 0 {
			evens = append(evens, value)
		}
	}

	fmt.Println(evens)


	// Problem 22: Odd numbers filter করো

	odds := []int{}

	for _, value := range slice8 {
		if value%2 != 0 {
			odds = append(odds, value)
		}
	}


	// Problem 23: Slice এ কোনো value আছে কিনা check করো
	fmt.Println(odds)	

	contains(slice1, 2)

	// Problem 24: Duplicate remove করো
		//  not found

	// Problem 25: Slice sort করো
	slice9 := []int{10, 50, 20, 80, 5}

	/// using sort package
	// sort.Ints(slice9)
	strings := []string{ "rahim", "jobbar", "arnob"}
	sort.Strings(strings)
	fmt.Println(strings)


	// using slices package
	slices.Sort(slice9)
	fmt.Println(slice9, slices.IsSorted(slice9))
	

	fruits  := []string{"peach", "banana", "kiwi"}

	// using slices package
	slices.Sort(fruits)
	fmt.Println(fruits, slices.IsSorted(fruits))




	slice10 := []int{10, 50, 20, 80, 5}
	slice11 := []int{3, 4, 12, 2, 11}

	//  appending
	fmt.Println(slice10,slice11)
	var slice12 []int
	slice12 = append(slice10, slice11...)
	fmt.Println(slice12)
	// deleting   

	// Problem 27: Slice থেকে element delete করো
	index := 2
	slice10 = append(slice10[:index], slice10[index+1:]...)
	fmt.Println(slice10)


	// Problem 28: Slice এর শুরুতে element insert করো
	insertEl  := 99
	slice10 = append([]int{insertEl}, slice10...)
	fmt.Println(slice10)




	// Problem 29: নির্দিষ্ট index এ element insert করো
	slice11 = []int{3, 4, 12, 2, 11}
	targetIndex := 2
	insertElement := 10

	left  := append([]int{}, slice11[:targetIndex]...)
	left = append(left, insertElement)

	right := slice11[targetIndex:]

	slice11 = append(left, right...)
	fmt.Println(slice11)
	


	// Problem 30: Two slices merge করো
	merged := append(slice1, slice2...)
	sorting([]string{"arnob", "jobbar", "rahim"})
	fmt.Println(merged)



}
	// fmt.Println(slice9)

	// string sorting