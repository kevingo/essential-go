package main

import "fmt"

func main() {
	var x [58]string

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	println("Array length: ", len(x))

	for i, v := range x {
		fmt.Println("Decimal: ", (i + 65), ", Value:", v)
	}

	var y [3]int
	fmt.Println(y) // [0,0,0]

	var z = [3]int{1, 2, 3}
	fmt.Println(z) // [1,2,3]

	a := [3]int{1, 2, 3}
	fmt.Println(a) // [1,2,3]

	b := [3]int{1}
	fmt.Println(b) // [1,0,0]

	c := [...]int{1, 2, 3}
	fmt.Println(c) // [1,2,3]

	println("\nExample2\n")
	arr2 := make([]int, 10)
	for i, v := range arr2 {
		fmt.Println("index: ", i, " value: ", v)
	}

	var name [2]string
	name[0] = "Kevin"
	name[1] = "John"

	fmt.Println(name[0])
	fmt.Println(name[1])

	bool_array := []bool{true, false, true, true}
	fmt.Println(bool_array)

	float_array := []float32{1.0, 2.0, 3.0, 1.8, 5.5}
	fmt.Println(float_array)

	float_array_2 := []float64{1.1, 1.0, 2.0, 2.2}
	fmt.Println(float_array_2)

	uint := []uint8{0, 1, 2, 3}
	fmt.Println(uint)

	int_array := []int8{-1, 0, 1}
	fmt.Println(int_array)
}
