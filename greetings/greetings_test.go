package greetings

import (
 "testing"
)

func TestGreet(t *testing.T) {
 name := "Barani"
 want := "Hello Barani"
 msg, err := Greet(name)
 if msg != want || err != nil {
  t.Fatalf("TestGreet(%s): got %q/%v, want %q/nil", name, msg, err, want)
 }
}
