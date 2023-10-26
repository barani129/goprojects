package main

import "fmt"

func main() {
  s, i, f := "hello", 45, 10.20
  fmt.Printf("%v is of type %T\n", s, s)
  fmt.Printf("%v is of type %T\n", i, i)
  fmt.Printf("%v is of type %T\n", f, f)
}
