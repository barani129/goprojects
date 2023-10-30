package main

import (
	"fmt"
	"math/rand"
)


func main() {
	num := rand.Intn(250)
	fmt.Printf("the value of number is: %v\t", num)
	if num >= 0 && num <= 100 {
		fmt.Println("between 0 and 100")
	}else if num >= 100 && num <= 200 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 200 and 250")
	}

}
