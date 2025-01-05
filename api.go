package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationData struct {
  Count int `json:"count"`
  Next string `json:"next"`
  Previous string `json:"previous"`
  Results []struct{
    Name string `json:"name"`
    Url string `json:"url"`
  } `json:"results"`
}

func getNextUrl (c *config) (string, error){

  //TODO: handle being on the last page by returning an error

  if c.next == ""{
    return "https://pokeapi.co/api/v1/location-area/", nil
  } else {
    return c.next, nil
  }

}

func getPrevUrl (c *config) (string, error){

  if c.previous == "" {
    return "", fmt.Errorf("you're on the first page")
  } else {
    return c.previous, nil
  }

}

func getLocation(c *config, next bool) ([]string, error) {

  var url string

  var err error

  switch next {
  case true:
    url, err = getNextUrl(c)
  case false:
    url, err = getPrevUrl(c)
  }
  
  if err != nil {
    return []string{}, err
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
  
  locationData := locationData{}
  err = json.Unmarshal(body, &locationData)

  if err != nil {
    return []string{}, err
  }

  locationList := []string{}

  for _, location := range locationData.Results {
    locationList = append(locationList, location.Name)
  }

  c.next = locationData.Next
  c.previous = locationData.Previous
  
  return locationList , nil
  
}
