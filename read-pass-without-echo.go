package main

import (
"fmt"
"log"
"os"
"github.com/x/term"
)

func main() {
fmt.Printf("Please enter your password: ")
data, err := term.ReadPassword(os.Stdin.Fd())
if err != nil {
log.Fatal(err)
}
fmt.Println(data)
}
