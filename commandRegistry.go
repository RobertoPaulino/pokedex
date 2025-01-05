package main

import (
  "fmt"
  "os"
)

func commandList() map[string]cliCommand {
   
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
    "map": {
      name: "map",
      description: "Displays the name of 20 locations in the Pokemon world, subsequent calls display the next 20 locations",
      callback: commandMap,
    },
  } 
}

func commandExit() error {
  fmt.Println("Closing the Pokedex... Goodbye!")

  os.Exit(0)

  return nil
}

func commandHelp() error {
  fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")

  commands := commandList()

  if len(commands) == 0 {
    return fmt.Errorf("No Command found")
  }

  for _, command := range commands {
    fmt.Printf("%v: %v\n", command.name, command.description)
  }

  return nil
}


func commandMap() error {
  
  locations, err := getLocation() 

  if err != nil {
    return err
  }
  println(locations)

  return nil

}
