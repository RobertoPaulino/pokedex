package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/robertopaulino/pokedex/internal/pokecache"
	"github.com/robertopaulino/pokedex/internal/pokedex"
)

func CommandList() map[string]cliCommand {

  mapConfig := config{
    next: "",
    previous: "",
  }
  
  mapCache := pokecache.NewCache(5 * time.Second)
  exploreCache := pokecache.NewCache(5 * time.Second)

  pokedex := pokedex.GetPokedex()

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
      Cache: mapCache,
    },
    "mapb": {
      Name: "map back",
      description: "Displays the name of the previous 20 locations in the pokemon world, if user is on the first page it will just let the user know they are on the first page",
      Callback: commandMapBack,
      Config: &mapConfig,
      Cache: mapCache,
    },
    "explore": {
      Name: "explore",
      description: "Displays pokemon found in the area, needs a ID or location name as a parameter.",
      Callback: commandExplore,
      Cache: exploreCache,
      Parameter: []string{},
    },
    "catch": {
      Name: "catch",
      description: "Attempts to catch pokemon",
      Callback: commandCatch,
      Parameter: []string{},
      Pokedex: pokedex,
    },
  } 
}

func commandCatch(config *config, cache *pokecache.Cache, parameters []string, pokedex *pokedex.Pokedex) error { 
  if len(parameters) < 1 {
    return fmt.Errorf("Not enough parameters")
  } 

  if len(parameters) > 1 {
    return fmt.Errorf("Too many parameters")
  }

  err := catch(parameters[0], pokedex)
  if err != nil {
    return err
  }
  
  return nil

}

func commandExplore(config *config, cache *pokecache.Cache, parameters []string, pokedex *pokedex.Pokedex) error {
  
  if len(parameters) < 1 {
    return fmt.Errorf("Not enough parameters")
  } 
  
  if len(parameters) > 1 {
    return fmt.Errorf("Too many parameters")
  }
  
  fmt.Printf("Exploring %v...\nFound Pokemon:\n", parameters[0])
  
  pokemonList, err := getPokemonList(cache, parameters[0])

  if err != nil {
    return err
  }
  for _, pokemon := range pokemonList {
    fmt.Printf(" - %v\n", pokemon)
  } 

  return nil

}

func commandExit(config *config, cache *pokecache.Cache, parameters []string, pokedex *pokedex.Pokedex) error {
  fmt.Println("Closing the Pokedex... Goodbye!")

  os.Exit(0)

  return nil
}

func commandHelp(config *config, cache *pokecache.Cache, parameters []string, pokedex *pokedex.Pokedex) error {
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


func commandMap(config *config, cache *pokecache.Cache, parameters []string, pokedex *pokedex.Pokedex) error {
  
  locations, err := getLocation(config, cache, true)

  if err != nil {
    return err
  }
  for _, location := range locations {
    fmt.Printf("%v\n", location)
  } 

  return nil

}

func commandMapBack(config *config, cache *pokecache.Cache, parameters []string, pokedex *pokedex.Pokedex) error {
  locations, err := getLocation(config, cache, false)

  if err != nil {
    return err
  }
  for _, location := range locations {
    fmt.Printf("%v\n", location)
  } 

  return nil
}
