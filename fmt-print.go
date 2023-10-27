package main

import "fmt"

func main() {
	//printf prints the output to stdout
	fmt.Printf("Binary: %b\\%b\n", 4, 5)
	//%% produces a single %d
	fmt.Printf("value is: %d %%\n", 50)
	//producing error
	fmt.Printf("Binary: %b/%b\n", 4)
}
