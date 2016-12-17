package main

import (
	"fmt"
)

type people struct {
	name string
	age  int
}

func main() {

	// 1. Struct
	p2 := people{"Kevin", 30}
	fmt.Println(p2)

	p2.rename("Mina")
	fmt.Println(p2)

	p2.rename2("Kevin")
	fmt.Println(p2)

	// 2. Map
	var m = make(map[string]int)
	m["age"] = 10
}

// receiver is a pointer type, it will change p
func (p *people) rename(name string) {
	p.name = name
}

// receiver is a value type, it will not change p
func (p people) rename2(name string) {
	p.name = name
}
