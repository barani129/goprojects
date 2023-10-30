package main

import (
	"fmt"
	"math/rand"
)


func main() {
	x := rand.Intn(250)
	fmt.Printf("the value of number is: %v\t", x)
	switch {
	case x <= 100:
	fmt.Println("between 0 and 100")
	case x <= 200:
	fmt.Println("between 101 and 200")
	case x <= 250:
	fmt.Println("betwen 201 and 250")
	default:
	fmt.Println("more than 250")
	}
}

func init() {
	fmt.Println("This is the initialization function")
}
