package main

import "fmt"

func visit(num []int, callback func(int)) {
	for _, n := range num {
		callback(n)
	}
}

func filter(num []int, c func(int) bool) []int {
	xs := []int{}
	for _, n := range num {
		if c(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func main() {
	println("\nExample 1\n")
	visit([]int{1, 2, 3, 4, 5}, func(n int) {
		println(n)
	})

	println("\nExample 2\n")
	ss := filter([]int{1, 2, 3, 4, 5}, func(n int) bool {
		return n > 1
	})

	fmt.Println(ss)
}

// callback: passing a func as an argument