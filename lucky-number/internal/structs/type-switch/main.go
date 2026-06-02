package main

import "fmt"

func main() {
	var value any

	
	value = 10
	// ugly code
	
	// if i, ok := v.(int); ok {
	// 	fmt.Println(i)
	// } else if s, ok := v.(string); ok {
	// 	fmt.Println(s)
	// }
	
	switch v := value.(type) {
		case int:
			fmt.Println("int", v)

		case string:
			fmt.Println("string", v)

		case bool:
			fmt.Println("bool", v)

		default:
			fmt.Println("unknown")
}
}

