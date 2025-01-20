package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) FetchPokemonByName(name *string) (PokemonData, error) {
	if name == nil {
		return PokemonData{}, fmt.Errorf("Missing Pokemon Name")
	}

	pageURL := c.baseURL + "/pokemon/" + *name
	pokemonData := PokemonData{}
	if data, found := c.cache.Get(pageURL); found {
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			fmt.Println(err)
			return PokemonData{}, err
		}
		return pokemonData, nil
	}
	req, err := http.NewRequest("GET", pageURL, nil)
	if err != nil {
		return PokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(pageURL, dat)
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return PokemonData{}, err
	}

	return pokemonData, nil
}
