package main

import "fmt"

type ByteSize int

const (
	_  =	iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Println(MB)
}
