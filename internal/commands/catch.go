package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/robertopaulino/pokedex/internal/pokedex"
	"github.com/robertopaulino/pokedex/internal/pokemon"
)

type PokemonData struct {
  Height int `json:"height"`
  Weight int `json:"weight"`
  BaseExperience int `json:"base_Experience"`
  Stats []struct {
    BaseStat int `json:"base_stat"`
    Stat struct {
      StatName string `json:"name"`
    }`json:"stat"`
  } `json:"stats"`
  Name string `json:"name"`
  Types []struct{
    Type struct{
      Name string `json:"name"`
    } `json:"type"`
  } `json:"types"`
}

func catch(name string, dex *pokedex.Pokedex) error {
 
  
  url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", name)

  pokemonData := PokemonData{}
  

  res, err := http.Get(url)
  if res.StatusCode == 404 {
    return fmt.Errorf("That pokemon does not exist!\n")
  }
  if err != nil {
    return err
  }

  body, err := io.ReadAll(res.Body)
  defer res.Body.Close()
  if res.StatusCode > 299 {
    return fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
  }
  if err != nil {
    return err
  }

  err = json.Unmarshal(body, &pokemonData)
  if err != nil {
    return err
  }
  
  isCaught := catchCalc(pokemonData.BaseExperience)
  
  pokemonStats := getStats(pokemonData)
  pokemonTypes := getTypes(pokemonData)
  if isCaught {
    newPokemon := pokemon.Pokemon{
      Name: pokemonData.Name,
      Height: pokemonData.Height,
      Weight: pokemonData.Weight,
      Types: pokemonTypes,
      Stats: pokemonStats,
    }    
    //TODO -> catching is causing panic due to adding nil entry to map somehow
    fmt.Printf("\n%v was caught!\n", newPokemon.Name)
    dex.AddPokemon(newPokemon)
  } else {
    fmt.Printf("\n%v ran awy...\n", pokemonData.Name)
  }

  return nil

}

func catchCalc(baseExperience int) bool {

  chance := baseExperience / 20
  if chance <= 0 {
    return true 
  }

  ceiling := rand.IntN(chance)

  if ceiling != 0 {
    return false
  }

  return true

}

func getStats(pokemonData PokemonData) pokemon.PokemonStats{
  statsData := pokemonData.Stats
  pokemonStats := pokemon.PokemonStats{} 
  for _, stat := range statsData{
    switch stat.Stat.StatName{
    case "hp":
      pokemonStats.Hp = stat.BaseStat
    case "attack":
      pokemonStats.Attack = stat.BaseStat
    case "defense":
      pokemonStats.Defense = stat.BaseStat
    case "special-attack":
      pokemonStats.SpecialAttack = stat.BaseStat
    case "special-defense":
      pokemonStats.SpecialDefense = stat.BaseStat
    case "speed":
      pokemonStats.Speed = stat.BaseStat
    default:
      fmt.Printf("Warning, uknown stat: %v\n", stat.Stat.StatName)
  }
  }
  
  return pokemonStats
}

func getTypes(pokemonData PokemonData) []string{
  
  types := []string{}

  for i := 0; i < len(pokemonData.Types) - 1; i++ {
    types[i] = pokemonData.Types[i].Type.Name
  } 

  return types
}
