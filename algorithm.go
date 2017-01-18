package main

import (
	algo "./algorithm"
	"fmt"
)

func main() {
	searchTarget := 1112
	initSearchSpace := []int{1, 2, 6, 10, 33, 55, 1112}

	found := algo.BinarySearch(searchTarget, initSearchSpace)
	fmt.Printf("Search %d in %d ? %t", searchTarget, initSearchSpace, found)
}
