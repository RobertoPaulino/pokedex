package main

import (
  "fmt"
  "strings"
)
func main() {
  fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {

  lower := strings.ToLower(text)
  trimmed := strings.Trim(lower, " ")
  split := strings.Split(trimmed, " ")

  return split
}
