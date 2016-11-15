package main

import "fmt"

func main() {
	switch "a" {
	case "a":
		fmt.Println("a")
		break
	case "b":
		fmt.Println("b")
		break
	default:
		fmt.Println("default")
	}

	switch "b" {
	case "a":
		fmt.Println("a")
		break
	case "b":
		fallthrough
	case "c":
		fmt.Println("b and c")
		break
	default:
		fmt.Println("default")
	}

	switch "c" {
	case "a":
		fmt.Println("a")
		break
	case "b", "c":
		fmt.Println("b and c")
		break
	default:
		fmt.Println("default")
	}

	var name = "Kevin"

	switch {
	case name == "a":
		fmt.Println("a")
		break
	case name == "b":
		fmt.Println("b")
		break
	default:
		fmt.Println("nothing match")
	}

}