package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type area struct {
  Name string `json:"name"`
  Url string `json:"url"`
}
  
type locationData struct {
  Count int `json:"count"`
  Next string `json:"next"`
  Previous string `json:"previous"`
  Results []map[int]area `json:"results"`
}

func getLocation() ([]string, error) {
  res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
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
  
  locationData := locationData{}
  err = json.Unmarshal(body, &locationData)

  if err != nil {
    fmt.Print("I ran!!!")
    return []string{}, err
  }

  locationList := []string{}

  for i, location := range locationData.Results {
    locationList = append(locationList, location[i].Name)
  }

  return locationList , nil

  
}
