package main

import (
  "fmt"
  "strings"
  "bufio"
  "os"
)

type cliCommand struct {
  name        string
  description string
  callback    func() error
}


func commandList() map[string]cliCommand{
   
  return map[string]cliCommand{

    "exit": {
      name: "exit",
      description: "Exit the Pokedex",
      callback: commandExit,
    },
    "help": {
      name: "help",
      description: "Displays a help message",
      callback: commandHelp,
    },
  } 
}


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
      err := commandInfo.callback()

      if err != nil {
        fmt.Printf("Error: %v \n input: %v \n running command %v\n", err, cleanText[0], commandInfo.name)
      }
      continue
    }

    fmt.Println("Unknown command")
  }
}


func cleanInput(text string) []string {

  lower := strings.ToLower(text)
  trimmed := strings.Trim(lower, " ")
  split := strings.Split(trimmed, " ")

  return split
}

func commandExit() error {
  fmt.Println("Closing the Pokedex... Goodbye!")

  os.Exit(0)

  return nil
}

func commandHelp() error {
  fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

  commands := commandList()

  for _, command := range commands {
    fmt.Printf("%v: %v\n", command.name, command.description)
  }

  return nil
}
