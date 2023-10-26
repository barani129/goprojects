package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, s := range os.Environ() {
		kv := strings.SplitN(s, "=", 2)
		fmt.Printf("Key: %q Value: %q\n", kv[0], kv[1])
	}
}
