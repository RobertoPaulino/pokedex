package main

import (
  "fmt"
  "os"
)

func commandList() map[string]cliCommand {

  mapConfig := config{
    next: "",
    previous: "",
  }

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
      config: &mapConfig,
    },
    "mapb": {
      name: "map back",
      description: "Displays the name of the previous 20 locations in the pokemon world, if user is on the first page it will just let the user know they are on the first page",
      callback: commandMapBack,
      config: &mapConfig,
    },
  } 
}

func commandExit(config *config) error {
  fmt.Println("Closing the Pokedex... Goodbye!")

  os.Exit(0)

  return nil
}

func commandHelp(config *config) error {
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


func commandMap(config *config) error {
  
  locations, err := getLocation(config, true)

  if err != nil {
    return err
  }
  for _, location := range locations {
    fmt.Printf("%v\n", location)
  } 

  return nil

}

func commandMapBack(config *config) error {
  locations, err := getLocation(config, false)

  if err != nil {
    return err
  }
  for _, location := range locations {
    fmt.Printf("%v\n", location)
  } 

  return nil
}
