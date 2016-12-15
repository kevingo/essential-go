package main

import (
	"fmt"
)

type Key struct {
	Path, Country string
}

func main() {
	hits := make(map[Key]int)

	hits[Key{"/", "tw"}]++
	hits[Key{"/doc", "us"}]++
	hits[Key{"/", "tw"}]++

	// fmt.Println(hits)

	fmt.Println("TW visit / count: ", hits[Key{"/", "tw"}])
	fmt.Println("TW visit /doc count: ", hits[Key{"/doc", "tw"}])
	fmt.Println("US visit / count: ", hits[Key{"/", "us"}])
	fmt.Println("US visit /doc count: ", hits[Key{"/doc", "us"}])
}
