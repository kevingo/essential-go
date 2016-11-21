package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	age := 30
	fmt.Println(&age) // print the address of age

	address := &age
	value := *address // pointer to the address to get the value

	fmt.Println(address, value) // print the address and the value of age

	*address = 10 // change the value by using pointer to the address
	fmt.Println(age) // age should be 10

	println("\nExample2\n")

	p1 := Point{10, 20}
	p2 := &Point{10, 20}

	fmt.Println(p1)
	fmt.Println(p2)
}