package greetings

import "fmt"
import "errors"

func Greet(name string) (string, error) {
  if name == "" {
  return "", errors.New("Empty Name")
  }
  message := fmt.Sprintf("Hello %v", name)
  return message, nil
}
