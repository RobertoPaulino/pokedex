package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  scanner := bufio.NewScanner(reader)

  commands := commandList()
  var text string

  for {
    fmt.Printf("Pokedex >")
    scanner.Scan()
    text = scanner.Text()

    cleanText := cleanInput(text)

    commandInfo, ok := commands[cleanText[0]] 

    if ok {
      err := commandInfo.callback(commandInfo.config)

      if err != nil {
        fmt.Printf("Error: %v \n input: %v \n running command %v\n", err, cleanText[0], commandInfo.name)
      }
      continue
    }

    fmt.Println("Unknown command")
  }
}
