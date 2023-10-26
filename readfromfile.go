package main

import (
  "fmt"
  "os"
)

func main() {
   data, err := os.ReadFile("/Users/barani/Downloads/dummy.txt")
   if err == nil {
    mydata := string(data)
    fmt.Println(mydata)
   }else {
    fmt.Println(err)
    }
}

