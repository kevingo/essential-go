package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5} // A slice type is a reference type, it has no specified length.
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
}
