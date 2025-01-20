package main

import (
	"math/rand"
	pokeapi "pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	caught map[string]pokeapi.PokemonData
}

func NewPokedex() *Pokedex {
	pokedex := &Pokedex{
		caught: make(map[string]pokeapi.PokemonData),
	}
	return pokedex
}

func (p *Pokedex) calculateCatchProb(baseExp int) float64 {
	minProb, maxProb := 0.2, 0.9
	minBaseExp, maxBaseExp := 50, 300
	return maxProb - float64(((baseExp-minBaseExp)*(int(maxProb)-int(minProb)))/(maxBaseExp-minBaseExp))
}

func (p *Pokedex) AttemptCatch(pokemonData pokeapi.PokemonData) bool {
	probability := p.calculateCatchProb(pokemonData.BaseExperience)
	randNum := rand.Float64()
	return randNum < probability
}

func (p *Pokedex) AddToPokedex(pokemonName string, pokemonData pokeapi.PokemonData) {
	p.caught[pokemonName] = pokemonData
}
