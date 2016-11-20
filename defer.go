package main

func hello() {
	println("hello")
}

func world() {
	println("world")
}

func main() {
	defer world()
	hello()
}