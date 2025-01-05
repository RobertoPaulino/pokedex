package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/robertopaulino/pokedex/internal/commands"
	"github.com/robertopaulino/pokedex/internal/utils"
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  scanner := bufio.NewScanner(reader)

  commands := commands.CommandList()
  var text string

  for {
    fmt.Printf("Pokedex >")
    scanner.Scan()
    text = scanner.Text()

    cleanText := utils.CleanInput(text)

    commandInfo, ok := commands[cleanText[0]] 

    if ok {
      err := commandInfo.Callback(commandInfo.Config)

      if err != nil {
        fmt.Printf("Error: %v \n input: %v \n running command %v\n", err, cleanText[0], commandInfo.Name)
      }
      continue
    }

    fmt.Println("Unknown command")
  }
}
