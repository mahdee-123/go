package main

import (
	"cmp"
	"fmt"
	"slices"
)

func sorting(slice []string) {
	

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a),len(b))
	}

	slices.SortFunc(slice, lenCmp)
	fmt.Println(slice)
}

