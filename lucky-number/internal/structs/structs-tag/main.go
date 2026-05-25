package main


import (
	"encoding/json"
	"fmt"
)
// struct tag = meta data
type User struct {
	Name string `json:"name,omitempty"`
	Age int 		`json:"age,omitempty"`
	Password string `json:"-"`
}


type Product struct {
	Id int 		`db:"id"`
	Price int   `db:"price"`
}


func main() {
	user := User{Name: "Arnob", Age: 20}
	data , _ := json.Marshal(user)
	fmt.Println(string(data))


	// 1. omitempty
	// Empty value হলে field skip করবে।
	user1 := User{Name: "user1", Password: "12498349"}
	data1 , _ := json.Marshal(user1)
	fmt.Println(string(data1))

	// 2. Ignore Fiel
	
}