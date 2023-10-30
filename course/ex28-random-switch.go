package main

import (
"fmt"
"math/rand"
)


func main() {
x := rand.Intn(10)
y := rand.Intn(10)
fmt.Printf("x value is: %v\t and y value is: %v\n", x, y)
switch {
case x < 4 && y < 4:
fmt.Println("both x and y are less than 4")
case x > 6 && y > 6:
fmt.Println("both x and y are greater than 6")
case x >= 4 && x <= 6:
fmt.Println("x is from 4 to 6")
case y != 5:
fmt.Println("y is not 5")
default:
fmt.Println("unknown values")
}
}
