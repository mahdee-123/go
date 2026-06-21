package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    // JSON data - যেকোনো type এর mixture হতে পারে
    jsonData := `{
        "name": "রহিম",
        "age": 25,
        "salary": 50000.50,
        "isActive": true,
        "skills": ["Go", "Python", "JavaScript"]
    }`
    
    // interface{} ব্যবহার করে যেকোনো JSON parse করা যায়
    var data interface{}
    json.Unmarshal([]byte(jsonData), &data)
    
    fmt.Println(data)
}