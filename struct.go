package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	p1 := person{name: "Kevin", age: 30}
	show(p1)

	p2 := person{"Mina", 31}
	show(p2)

	p3 := person{name: "John"}
	show(p3)

	p4 := person{"Fu", 10}
	show(p4)
	change(&p4)
	show(p4)

}

func show(p person) {
	fmt.Println(p.name, p.age)
}

func change(p *person) {
	p.name = p.name + "_new"
}