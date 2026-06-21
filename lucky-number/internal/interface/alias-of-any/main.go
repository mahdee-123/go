package main

import "fmt"

func PrintAnything(value interface{}) {
    fmt.Println("Value:", value)
}

func main() {
    PrintAnything(42)              // ✅ int
    PrintAnything(3.14)            // ✅ float64
    PrintAnything("Hello")         // ✅ string
    PrintAnything(true)            // ✅ bool
    PrintAnything([]int{1, 2, 3})  // ✅ slice
    PrintAnything(map[string]int{}) // ✅ map
    
    // সব কিছুই কাজ করে!
}