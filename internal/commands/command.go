package commands

import "github.com/robertopaulino/pokedex/internal/pokecache"


type cliCommand struct {
	Name        string
	description string
	Callback    func(*config) error
	Config      *config
	cache *pokecache.Cache
}

type config struct {
	next     string
	previous string
}
