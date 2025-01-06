package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/robertopaulino/pokedex/internal/pokecache"
)

func CommandList() map[string]cliCommand {

  mapConfig := config{
    next: "",
    previous: "",
  }
  
  mapCache := pokecache.NewCache(5 * time.Second)

  return map[string]cliCommand{

    "exit": {
      Name: "exit",
      description: "Exit the Pokedex",
      Callback: commandExit,
    },
    "help": {
      Name: "help",
      description: "Displays a help message",
      Callback: commandHelp,
    },
    "map": {
      Name: "map",
      description: "Displays the name of 20 locations in the Pokemon world, subsequent calls display the next 20 locations",
      Callback: commandMap,
      Config: &mapConfig,
      cache: mapCache,
    },
    "mapb": {
      Name: "map back",
      description: "Displays the name of the previous 20 locations in the pokemon world, if user is on the first page it will just let the user know they are on the first page",
      Callback: commandMapBack,
      Config: &mapConfig,
      cache: mapCache,
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

  commands := CommandList()

  if len(commands) == 0 {
    return fmt.Errorf("No Command found")
  }

  for _, command := range commands {
    fmt.Printf("%v: %v\n", command.Name, command.description)
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
