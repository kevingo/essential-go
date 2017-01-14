package main

import (
	"fmt"
)

func main() {
	s := make(chan string) // initial a string channel

	// initial another goroutine
	go func() {
		s <- "hello" // s is ready
	}()

	val := <-s // val is waiting s to ready

	fmt.Println(val)

	/* Example 2 */
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{6, 7, 8, 9, 10}

	c1 := make(chan int)
	c2 := make(chan int)
	go sum(s1, c1)
	go sum(s2, c2)

	ans1 := <-c1
	ans2 := <-c2

	fmt.Println(ans1, ans2)

	/* Example 3 Buffered Channel */

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 this line will cause fatal error
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}
