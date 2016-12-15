package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[int]string)
	var keys []int

	m[0] = "Kevin"
	m[1] = "John"
	m[2] = "Mina"

	for k, _ := range m {
		fmt.Println("Key:", k, " Value: ", m[k])
	}

	fmt.Println()

	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println("Key:", k, " Value: ", m[k])
	}

}
