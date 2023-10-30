package main

import "fmt"

func main() {
y := 5
for {
fmt.Printf("The value of y is: %v\n", y)
if y > 30 {
break
}
y++
}
}
