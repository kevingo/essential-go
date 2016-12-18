package main

import (
	"fmt"
)

type people struct {
	name string
	age  int
}

type Map map[string]int

func main() {

	// 1. Struct
	p2 := people{"Kevin", 30}
	fmt.Println(p2)

	p2.rename("Mina")
	fmt.Println(p2)

	p2.rename2("Kevin")
	fmt.Println(p2)

	// 2. map
	m := Map{"age": 10}
	fmt.Println(m)

	m.add("salary", 10000)
	fmt.Println(m)
}

// receiver is a pointer type, it will change p
func (p *people) rename(name string) {
	p.name = name
}

// receiver is a value type, it will not change p
func (p people) rename2(name string) {
	p.name = name
}

func (m Map) add(key string, value int) {
	m[key] = value
}
