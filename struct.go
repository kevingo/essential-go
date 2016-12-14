package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
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
	change2(p4)
	show(p4)

	p5 := new(person)
	p5.name = "Kevin"
	p5.goHome()
	p5.goHome2()
}

func show(p person) {
	fmt.Println(p.name, p.age)
}

func change(p *person) {
	p.name = p.name + "_new"
}

func change2(p person) {
	p.name = p.name + "_new?"
}

// receiver
func (p *person) goHome() {
	p.name = "Mina"
	fmt.Println(p.name + ", let's go home")
}

func (p person) goHome2() {
	fmt.Println(p.name + ", let's go home2")
}
