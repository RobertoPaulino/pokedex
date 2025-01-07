package commands

import (
	"github.com/robertopaulino/pokedex/internal/pokecache"
	"github.com/robertopaulino/pokedex/internal/pokedex"
)

type cliCommand struct {
	Name        string
	description string
	Callback    func(*config, *pokecache.Cache, []string, *pokedex.Pokedex) error
	Config      *config
	Cache *pokecache.Cache
  Parameter []string
  Pokedex *pokedex.Pokedex 
}

type config struct {
	next     string
	previous string
}
