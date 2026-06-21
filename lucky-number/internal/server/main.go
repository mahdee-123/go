package main

import (
	"fmt"
	"net/http"
)


func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world from server!")

}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "all users")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/users", GetUsers)
	http.ListenAndServe(":8080", nil)
}