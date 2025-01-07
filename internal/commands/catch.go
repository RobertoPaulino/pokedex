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
      Name string `json:"name"`
    }`json:"stat"`
  } `json:"stats"`
  name string
  types []string
}

func catch(name string, dex *pokedex.Pokedex) error {
 

  url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", name)

  pokemonData := PokemonData{}


  res, err := http.Get(url)
  if res.StatusCode == 404 {
    return fmt.Errorf("That pokemon does not exist!")
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

  if isCaught {
    // TODO finish this so we can move pokemonData -> pokemon and store it on the dex
    newPokemon := pokemon.Pokemon{
      //Name: ,
    }
    dex.AddPokemon(newPokemon)
  }

  return nil

}

func catchCalc(baseExperience int) bool {

  chance := baseExperience / 20
  if chance <= 0 {
    return true 
  }

  ceiling := rand.IntN(chance)
  fmt.Printf("capture chance 1 in %v\n", ceiling)

  if ceiling != 0 {
    return false
  }

  return true

}
