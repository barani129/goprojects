package main

import (
"fmt"
"math/rand"
)

func main() {
count := 0
for i := 0; i < 100; i++ {
if x := rand.Intn(5); x == 3 {
fmt.Printf("The value is: %v\n", x)
count++
}
}
fmt.Printf("The total count of 3s occurence is: %v\n", count)
}
