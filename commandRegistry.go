package main

import (
  "fmt"
  "os"
)

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
