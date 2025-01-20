package main

import (
	pokeapi "pokedexcli/internal/pokeapi"
	"pokedexcli/internal/pokecache"
	"time"
)

func main() {
	println("The Pokedex is STARTING !")
	println("@~@~@~@~@~@~@~@~@~@~@~@~@")
	config := Config{
		pokeapiClient: pokeapi.NewClient(pokecache.NewCache(10 * time.Second)),
		Pokedex:       NewPokedex(),
	}
	startRepl(&config)
}
