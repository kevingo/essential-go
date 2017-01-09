package main

import "fmt"

func main() {
	// A slice type is a reference type, it has no specified length.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	fmt.Println("Length: ", len(slice))
	fmt.Printf("Type: %T", slice)
	fmt.Println("\nFirst two elements: ", slice[0:2])
	fmt.Println("First two elements: ", slice[:2])
	fmt.Println("Last two elements: ", slice[len(slice)-2:len(slice)])
	fmt.Println("Third element: ", slice[2])
	fmt.Println("Build-in cap func cap(slice):", cap(slice))
	fmt.Println("abc"[0]) // decimal number of a, which is 97

	slice2 := make([]int, 5)
	fmt.Println("slice2: ", slice2)
	fmt.Println("len: ", len(slice2), " cap: ", cap(slice2))
	fmt.Printf("address of slice2 %p\n", &slice2)

	slice2 = append(slice2, 0)
	fmt.Println("len: ", len(slice2), " cap: ", cap(slice2))
	fmt.Printf("address of appended slice2 %p\n", &slice2)

	sub := slice2[2:4]
	fmt.Println("sub : ", sub)
	fmt.Println("sub len: ", len(sub), " cap: ", cap(sub))

	println("\nExample2\n")
	slice3 := make([]int, 5, 10)
	fmt.Println(slice3)
	fmt.Println("len: ", len(slice3), " cap: ", cap(slice3))
	fmt.Printf("address of slice3 %p\n", &slice3)
	slice3 = append(slice3, 0)
	fmt.Println(slice3)
	fmt.Printf("address of appended slice3 %p\n", &slice3)

	// append two slices
	println("\nExample3\n")
	hello := []string{"Hello"}
	world := []string{"world"}
	fmt.Println(append(hello, world...)) // note you have to add ...

	println("\nExample4\n")
	golang := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	partial := golang[0:2] // pointer to the golang slice
	fmt.Println(golang, partial)
	golang[0] = 'a'
	fmt.Println(golang, partial) // partial will also be changed to 'a'
	iamarray := [3]int{1, 2, 3}
	iamslice := iamarray[:] // a slice referencing the storage of iamarray
	fmt.Println(iamarray, iamslice)
	iamarray[0] = 0
	fmt.Println(iamarray, iamslice) // iamslice also be changed

	println("\nExample 5\n")
	a := []int{1, 2, 3, 4, 5}
	b := a[2:4]
	fmt.Println(a, cap(a))
	fmt.Println(b, cap(b))
	b = b[:cap(b)]
	fmt.Println(b, cap(b))

	println("\nExample 6\n")
	c := []int{1, 2, 3, 4, 5}
	fmt.Println("before ", c, cap(c))
	d := make([]int, len(c), (cap(c)+1)*2)
	copy(d, c)
	c = d
	fmt.Println("after ", c, cap(c))

	e := Filter(c, fn)
	fmt.Println("my e slice ", e)
}

func fn(s int) bool {
	if s%2 == 0 {
		return true
	}
	return false
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}
