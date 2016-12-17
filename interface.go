package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

type Backend struct {
	Human
	codingLevel int
}

type F2E struct {
	Human
	cssLevel        int
	javascriptLevel int
}

type RD interface {
	Coding()
}

func (p *Human) Eat() {
	fmt.Println("I can eat!!")
}

// implement Coding function means inherent RD
func (f2e *F2E) Coding() {
	fmt.Println("I write cool css and javascript!")
}

// implement Coding function means inherent RD
func (backend *Backend) Coding() {
	fmt.Println("I am backend and I can code!!")
}

func work(rd RD) {
	rd.Coding()
}

func main() {
	person := Human{name: "kevingo", age: 30}
	person.Eat()
	fmt.Println(person)

	abby := F2E{Human{name: "abby", age: 35}, 80, 90}
	fmt.Println(abby)
	kkgo := Backend{Human{name: "kevingo", age: 30}, 80}
	fmt.Println(kkgo)

	abby.Eat()
	kkgo.Eat()

	fmt.Println(kkgo.name)

	abby.Coding()
	kkgo.Coding()
	work(&abby)
	work(&kkgo)
}
