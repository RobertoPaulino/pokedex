package pokedex

import "github.com/robertopaulino/pokedex/internal/pokemon"

type Pokedex struct {
  pokedex map[string]pokemon.Pokemon
}


func GetPokedex() *Pokedex{
  var pokedex Pokedex
  return &pokedex
}

func (p *Pokedex) AddPokemon(data pokemon.Pokemon) {
  p.pokedex[data.Name] = data
}
