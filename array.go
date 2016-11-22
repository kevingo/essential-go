package main

import "fmt"

func main() {
    var x [58]string

    for i:= 65; i<=122; i++ {
        x[i-65] = string(i)
    }
    
    println("Array length: ", len(x))

    for i, v := range x {
        fmt.Println("Decimal: ", (i+65), ", Value:", v)
    }

    var y [3]int
    fmt.Println(y) // [0,0,0] 

    var z = [3]int{1, 2, 3}
    fmt.Println(z) // [1,2,3]

    a := [3]int{1, 2, 3}
    fmt.Println(a) // [1,2,3]

    b := [3]int{1}
    fmt.Println(b) // [1,0,0]

    c := [...]int{1,2,3}
    fmt.Println(c) // [1,2,3]

    println("\nExample2\n")
    arr2 := make([]int, 10)
    for i, v := range arr2 {
        fmt.Println("index: ", i, " value: ", v )
    }
}