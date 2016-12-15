package main

import (
	"fmt"
)

func main() {
	var m = make(map[string]int)
	fmt.Println(m)
	fmt.Printf("map type: %T\n", m)

	m["age"] = 30
	fmt.Println(m)
	fmt.Println(m["age"])
	fmt.Println(len(m))

	m["salary"] = 100
	fmt.Println(len(m))

	delete(m, "age")
	fmt.Println(m)

	v, ok := m["age"]  // value, exist or not
	fmt.Println(ok, v) //false

	m = map[string]int{"age": 30, "salary": 100}
	fmt.Println(m)

	m2 := map[string]string{}
	fmt.Println(m2)
	fmt.Println(m2 == nil) // false

	var m3 map[string]string // bad way to create a map
	fmt.Println(m3)
	fmt.Println(m3 == nil) // true

	m4 := map[int]string{
		1: "kevin",
		2: "john",
	}

	fmt.Println(m4)

	for i, v := range m4 {
		fmt.Println(i, " - ", v)
	}
}
