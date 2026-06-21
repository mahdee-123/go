// Because Textio is a glorified customer database, we have a lot of internal logic for sorting and dealing with customer names.

// Complete the getNameCounts function. It takes a slice of strings names and returns a nested map. The parent map's keys are all the unique first characters (see runes) of the names, the nested maps keys are all the names themselves, and the value is the count of each name.

// For example:
// {

// }
// billy
// billy
// bob
// joe

// Creates the following nested map:

// b: {
//     billy: 2,
//     bob: 1
// },
// j: {
//     joe: 1
// }

// Tips
// The test suite is not printing the map you're returning directly, but instead checking a few specific keys.
// You can convert a string into a slice of runes as follows:
// runes := []rune(word)
package main

import "fmt" 

func getNameCounts(names []string) map[rune]map[string]int {
	namesCount := make(map[rune]map[string]int)
	
	for _, name := range names {
		runes := []rune(name)
		firstChar := runes[0]
		fmt.Println(runes)
		// nil checking
		if namesCount[firstChar] == nil {
			namesCount[firstChar] = make(map[string]int)
		}

		namesCount[firstChar][name]++
	}
	return namesCount
}

func main() {
	names := []string {
		"billy",
		"billy",
		"bob",
		"joe",
	}

	fmt.Println(getNameCounts(names))
}