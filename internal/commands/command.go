package commands

import "github.com/robertopaulino/pokedex/internal/pokecache"


type cliCommand struct {
	Name        string
	description string
	Callback    func(*config, *pokecache.Cache) error
	Config      *config
	Cache *pokecache.Cache
}

type config struct {
	next     string
	previous string
}
