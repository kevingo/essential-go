package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    First string
    Last string
    Age int
    notExported int
}

func main() {
    p1 := Person{"Kevin", "Tsai", 20, 007}
    fmt.Println("p1 struct: ", p1)

    j, _ := json.Marshal(p1)
    fmt.Printf("Type %T\n", j)
    fmt.Println(string(j))
}