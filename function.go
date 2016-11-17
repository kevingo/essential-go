package main

import (
	"bytes"
	"fmt"
)

var name1 = "kevin"
var name2 = "john"

func main() {

	// test multiple return
	p1, p2 := f1()
	println(p1, p2)

	// multiple return with parameters
	p3, p4 := f2(name1, name2)
	println(p3, p4)

	p := f3(name1, name2)
	println(p)

	pp := f4(name1, name2)
	println(pp)

	// func expression: assign anonymous function to a variable
	hello := func() {
		println(name1, name2)
	}

	hello()

	// assign a func to a variable and execute it
	p5 := f5()
	p5()

	// execute a func immediately
	f5()()


}

func f1() (string, string) {
	return name1, name2
}

// multiple returns
func f2(n1, n2 string) (string, string) {
	return n1, n2
}

func f3(n1, n2 string) (s string) {
	s = fmt.Sprintf("%s %s", n1, n2)
	return s
}

// f4 is a variadic function, taking zero or more arguments
func f4(n ...string) string {
	var buffer bytes.Buffer

	for _, s := range n {
		buffer.WriteString(s + " ")
	}

	return buffer.String()
}

func f5() func() {
	return func() {
		println(name1, name2)
	}
}
