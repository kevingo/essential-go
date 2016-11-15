package main

import (
	"fmt"
)

func main() {
	age := 30
	fmt.Println(&age) // print the address of age

	address := &age
	value := *address // pointer to the address to get the value

	fmt.Println(address, value) // print the address and the value of age

	*address = 10 // change the value by using pointer to the address
	fmt.Println(age) // age should be 10
}