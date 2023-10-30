package main

import "fmt"

func main() {
xi := []int{42, 43, 44, 45, 46, 47}
for i, j := range xi {
fmt.Printf("The index is %v and its value: %v\n", i, j)
}
}
