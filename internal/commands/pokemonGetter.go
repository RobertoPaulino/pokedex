package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/robertopaulino/pokedex/internal/pokecache"
)

type encounterData struct {
  PokemonEncounters []struct{
    Pokemon struct{
      Name string `json:"name"`
    } `json:"pokemon"`
  }`json:"pokemon_encounters"`
}


func getPokemonList(cache *pokecache.Cache, locationName string) ([]string, error) {


  // TODO allow users to also search the area using ID instead
  url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", locationName)
  var err error
  
  encounterData := encounterData{}

  resCache, ok := cache.Get(url)
  if ok {
    err = json.Unmarshal(resCache, &encounterData)

    if err != nil {
      return []string{}, err
    }

    encounterList := []string{}

    for _, encounter := range encounterData.PokemonEncounters {
      encounterList = append(encounterList, encounter.Pokemon.Name)
    }

  
    return encounterList, nil
  }


  res, err := http.Get(url)
  if err != nil {
    return []string{}, fmt.Errorf("error requesting: %w", err)
  }
  body, err := io.ReadAll(res.Body)
  defer res.Body.Close()

  if res.StatusCode > 299 {
    return []string{}, fmt.Errorf("Unsucessful response code: %v", res.StatusCode)
  }

  if err != nil {
    return []string{}, fmt.Errorf("Unknown error: %w", err)
  } 

  err = json.Unmarshal(body, &encounterData)

  if err != nil {
    return []string{}, err
  }

  encounterList := []string{}

  for _, encounter := range encounterData.PokemonEncounters {
    encounterList = append(encounterList, encounter.Pokemon.Name)
  }

  cache.Add(url, body)
  
  return encounterList , nil
  
}
