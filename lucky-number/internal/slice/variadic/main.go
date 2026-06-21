package main

import "fmt"

func concat(strs ...string) string {
	final := ""
	for i := 0; i < len(strs); i++ {
		final += strs[i]
	}
	return final
}


func printStrings(strs ...string)  {
	for _, str := range strs {
		fmt.Println(str)
	}
}
func main() {
	fmt.Println(concat("hello", "this", "miss"))
	
	names := []string{"mahdee", "mahee", "arnob"}
	printStrings(names...)
}