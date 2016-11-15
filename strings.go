package main

import (
	"fmt"
)

func main() {
	for i := 50 ; i<= 140 ; i++ {
		// print string and ascii decimal
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	a := 'a' // a is a rune
	fmt.Println(a) // should be ascii decimal, that is 97 of rune a
	fmt.Printf("%T", a) // type of rune is int32

	foo := "a" // a is a string
	fmt.Println("\n" + foo) // should be a
	fmt.Printf("%T", foo) // string
}