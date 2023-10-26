package main

import (
  "fmt"
  "log"
  "example.com/greetings"
)

func main() {
  message, err := greetings.Greet("Venna")
  if err != nil {
   log.Fatal(err)
  }
 fmt.Println(message)
}
