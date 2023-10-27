package main

import (
	"fmt"
	"os"
)

func main() {
	data := "my data is here yoyo"
	newdata := []byte(data)
	if err := os.WriteFile("/Users/barani/Downloads/dummy.txt", newdata, 0644); err != nil {
		fmt.Println(err)
	}
}
