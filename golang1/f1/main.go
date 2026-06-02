package main
import (
	"fmt"
)


func main() {
	// 1. Create an int variable and print it.
	var x int = 34
	fmt.Println(x)

	// 2. int8 overflow
	var a int8 = 127
	a = a + 1
	fmt.Println(a)

	// 3. Store 255 in uint8.
	var c uint8 = 255
	fmt.Println(c)

	// 4. Float precision
	var d float32 = 12345678.9
	fmt.Println(d)

	// 5. Store true/false.
	var f bool = true
	fmt.Println(f)

	// 6.Type casting
	// var g bool = false 
	// var h  int8 = int8(g)
	// fmt.Println(h)

	var l int8 = 21 
	var m int16 = int16(l)
	fmt.Println(m)

	// 7. float → int
	var f1 float32 = 1.6
	fmt.Println(int(f1))

	// 8. int → float
	var i2 int8 = 10 
	fmt.Println(float32(i2))

	// 9. byte conversion
	
	var r2 byte = 'a'
	fmt.Println(r2)

	// 10. rune conversion
	var r1 rune = 'হ'
	fmt.Println(r1)
	
	// 11. String length
	str1 := "Hello"
	fmt.Println(str1[0])
	fmt.Println(string(str1[0]))
	fmt.Println(len(str1))
	
	
	// 12. Bangla string length
	
	str2 := "টি ছুঁয়ে আসবে"

	fmt.Println(str2)
	fmt.Println(str2[2])
	fmt.Println(len(str2))

	
	// 13. Rune count

	str3 := "বাংলা"
	fmt.Println([]rune(str3))
	fmt.Println(len([]rune(str3)))
	
	
	// 14. Loop over runes

	str4 := "বাংলা"
	for _, r := range str4 {
		fmt.Println(r)
		fmt.Println(string(r))
	}

	// 15. Print rune as character
	fmt.Printf("%c\n", 'a')

	// 16. Convert string to bytes
	str5 := "this is a string"
	fmt.Println([]byte(str5))

	// 17. Access first byte

	str6 := "this is a string"
	fmt.Println(str6[0])

	// 18. Print byte as char
	fmt.Printf("%c",str6[0])


	// 19. compare byte and rune 

	var x1 byte = 'a' // int8 
	var y1 rune = 'a' // int32
	fmt.Println(x1, y1)

	// 20. Emoji length
	var emoji string = "🫦"
	fmt.Println(len([]rune(emoji)))

	
	var nums [3]int = [3]int{1,34,4}
	fmt.Println(nums)
	
}


