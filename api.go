package main

import (
	"fmt"
	"io"
	"net/http"
)


func getLocation() (string, error) {
  res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
  if err != nil {
    return "", fmt.Errorf("error requesting: %w", err)
  }
  body, err := io.ReadAll(res.Body)
  defer res.Body.Close()

  if res.StatusCode > 299 {
    return "", fmt.Errorf("Unsucessful response code: %v", res.StatusCode)
  }

  if err != nil {
    return "", fmt.Errorf("Unknown error: %w", err)
  } 

  return string(body), nil
}
