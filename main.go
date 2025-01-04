package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)
func main() {
  reader := bufio.NewReader(os.Stdin)
  scanner := bufio.NewScanner(reader)

  var text string

  for {
    fmt.Printf("Pokedex >")
    scanner.Scan()
    text = scanner.Text()

    cleanText := cleanInput(text)

    fmt.Printf("Your command was: %v \n", cleanText[0])
  }
}

func cleanInput(text string) []string {

  lower := strings.ToLower(text)
  trimmed := strings.Trim(lower, " ")
  split := strings.Split(trimmed, " ")

  return split
}
