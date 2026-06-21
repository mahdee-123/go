package main

import (
	"fmt"
)

// Map দিয়ে Set বানাতে পারো!

// Set = যেখানে শুধু unique values থাকে (duplicates নেই)
//       কোনো order নেই
//       শুধু value আছে কিনা সেটা check করা যায়

// Implementation:
// - Map এর key = set এর values
// - Map এর value = bool (true/false)
// - Key আছে কিনা check করতে: indexing




func main() {
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}

	uniqueSet := map[int]bool{}
	
	fmt.Println(uniqueSet)
	for _, num := range numbers {
		uniqueSet[num] = true
	}

	fmt.Println(uniqueSet)

}