package main

import "fmt"

func main() {
m := map[string]int { 
"James": 42,
"Moneypenny": 32, 
}
//m1 := m["James"]
if age, ok := m["Jamess"]; ok {
fmt.Printf("The age of the user is: %v\n", age)
}else {
fmt.Println("The user is not found")
}
for k, v := range m {
fmt.Printf("The key is: %v and its value: %v\n", k, v)
}
}
