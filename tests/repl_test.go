package main

import (
  "testing"
)

func TestCleanInput(t *testing.T) {

  cases := []struct {
    input string
    expected []string
  }{
    {
      input: " hello world ",
      expected: []string{"hello", "world"},
    },

    {
      input: "HELLO world",
      expected: []string{"hello", "world"},
    },

    {
      input: "hello",
      expected: []string{"hello"},
    },

    {
      input: "WORLD",
      expected: []string{"world"},
    },

    {
      input: " hello world ",
      expected: []string{"hello", "world"},
    },
  }

  for _, c := range cases {
    actual := cleanInput(c.input)
    
    if len(c.expected) != len(actual) {
      t.Errorf("wrong length")
    }

    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]

      if word != expectedWord {
        t.Errorf("wrong word")
      }
    }
  
  }
}
