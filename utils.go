package main

import "strings"

func cleanInput(text string) []string {

  lower := strings.ToLower(text)
  trimmed := strings.Trim(lower, " ")
  split := strings.Split(trimmed, " ")

  return split
}

