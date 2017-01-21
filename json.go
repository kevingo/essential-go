package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

type MyUser struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	p1 := Person{"Kevin", "Tsai", 20, 007}
	fmt.Println("p1 struct: ", p1)

	// Example 1: Marchal return the json encode
	j, _ := json.Marshal(p1)
	fmt.Printf("Type %T\n", j)
	fmt.Println(string(j))

	// Example 2
	json.NewEncoder(os.Stdout).Encode(&MyUser{1, "kevingo", time.Now()})
}
