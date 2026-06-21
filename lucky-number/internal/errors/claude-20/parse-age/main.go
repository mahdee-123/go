// # parseAge(s string) (int, error) — strconv.Atoi ব্যবহার করে, fmt.Errorf দিয়ে %w wrap করো।

package main

import (
	"fmt"
	"strconv"
)


func parseAge(s string) (int, error) {
	num , err := strconv.Atoi(s)
	if err != nil {
		return 0, err 
	}
	return num, nil
}


func main() {

	fmt.Println(parseAge("34"))
	

}