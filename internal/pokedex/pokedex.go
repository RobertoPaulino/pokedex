package pokedex

import (
	"fmt"
	"sync"

	"github.com/robertopaulino/pokedex/internal/pokemon"
)

type Pokedex struct {
  pokedex map[string]pokemon.Pokemon
  mu sync.Mutex
}


func GetPokedex() *Pokedex{
  dex := Pokedex{
    pokedex: map[string]pokemon.Pokemon{},
    mu: sync.Mutex{},
  }
  return &dex
}

func (p *Pokedex) GetPokemon(name string) (pokemon.Pokemon, error) {

  pokemonInfo, ok := p.pokedex[name]

  if !ok {
    return pokemon.Pokemon{}, fmt.Errorf("That pokemon is not registered!")
  }

  return pokemonInfo, nil
}

func (p *Pokedex) AddPokemon(data pokemon.Pokemon) {

  p.mu.Lock()
  defer p.mu.Unlock()
  p.pokedex[data.Name] = data

}
