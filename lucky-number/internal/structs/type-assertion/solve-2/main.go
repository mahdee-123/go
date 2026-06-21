package main

import "fmt"

func ProcessValue(value interface{}) {
    // TODO: switch statement দিয়ে handle করো
    // সব type এর জন্য আলাদা handling
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

func main() {
    ProcessValue("Hello")
    ProcessValue(42)
    ProcessValue(3.14)
    ProcessValue(true)
}