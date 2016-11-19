package main

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	println(factorial(6)) // 6 * 5 * 4 * 3 * 2 * 1
}