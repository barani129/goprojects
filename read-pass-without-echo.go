package main

import (
"fmt"
"log"
"os"
"golang.org/x/term"
)

func main() {
fmt.Printf("Please enter your password: \n")
data, err := term.ReadPassword(int(os.Stdin.Fd()))
if err != nil {
log.Fatal(err)
}
fmt.Println("Reading your password")
fmt.Println(string(data))
}
